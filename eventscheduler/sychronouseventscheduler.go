package eventscheduler

import "fmt"

type SynchronousEventScheduler struct {
	Name         string
	eventChannel chan bool
	eventList    []Event
	stopCh       chan bool
}

func (s *SynchronousEventScheduler) AddEvent(ev Event) bool {
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

func (s *SynchronousEventScheduler) DeleteEvent(ev Event) bool {
	return s.DeleteEventById(ev.EvenId())
}

func (s *SynchronousEventScheduler) DeleteEventById(id Id) bool {
	for _, e := range s.eventList {
		if e.CompareWithId(id) {
			e.MarkEventActive(false)
			return true
		}
	}
	return false
}

func (s *SynchronousEventScheduler) popEvent() Event {
	for _, e := range s.eventList {
		if e.IsEventActive() {
			s.eventList = s.eventList[1:]
			return e
		}
	}
	return nil
}

func (s *SynchronousEventScheduler) ExecuteEventLoop() {
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

			e.Execute()
			e.MarkEventActive(false)

		case <-s.stopCh:
			fmt.Println("Existing sync event scheduler")
			return
		}
	}

}

func NewSyncEventScheduler(name string, stop chan bool) EventSchedulerInterface {
	return &SynchronousEventScheduler{
		Name:         name,
		eventChannel: make(chan bool),
		stopCh:       stop,
	}

}

var _ EventSchedulerInterface = &SynchronousEventScheduler{}
