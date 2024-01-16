package eventscheduler

import "fmt"

type ConcurrentEventScheduler struct {
	Name              string
	eventChannel      chan bool
	concurrentChannel chan bool
	eventList         []Event
	stopCh            chan bool
}

func (s *ConcurrentEventScheduler) AddEvent(ev Event) bool {
	for _, e := range s.eventList {
		if e.Compare(ev) {
			return false
		}
	}

	ev.MarkEventActive(true)
	eventId := len(s.eventList)
	ev.SetEventId(Id(eventId))
	s.eventList = append(s.eventList, ev)
	s.eventChannel <- true
	return true
}

func (s *ConcurrentEventScheduler) DeleteEvent(ev Event) bool {
	return s.DeleteEventById(ev.EvenId())
}

func (s *ConcurrentEventScheduler) DeleteEventById(id Id) bool {
	for _, e := range s.eventList {
		if e.CompareWithId(id) {
			e.MarkEventActive(false)
			return true
		}
	}
	return false
}

func (s *ConcurrentEventScheduler) popEvent() Event {
	for _, e := range s.eventList {
		if e.IsEventActive() {
			s.eventList = s.eventList[1:]
			return e
		}
	}
	return nil
}

func (s *ConcurrentEventScheduler) ExecuteEventLoop() {
	for {
		select {
		case <-s.eventChannel:
			e := s.popEvent()

			if e == nil {
				fmt.Println("Can not execute empty event")
				continue
			}

			if !e.IsEventActive() {
				fmt.Println("Can not execute inactive event ", e.EventName())
				continue
			}
			s.concurrentChannel <- true
			go func() {
				e.Execute()
				e.MarkEventActive(false)
				<-s.concurrentChannel
			}()

		case <-s.stopCh:
			fmt.Println("Existing sync event scheduler")
			return
		}
	}

}

func NewConcurrentEventScheduler(name string, stop chan bool, levelOfConc int) EventSchedulerInterface {
	return &ConcurrentEventScheduler{
		Name:              name,
		eventChannel:      make(chan bool),
		concurrentChannel: make(chan bool, levelOfConc),
		stopCh:            stop,
	}

}

var _ EventSchedulerInterface = &ConcurrentEventScheduler{}
