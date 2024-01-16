package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/arijit/eventscheduler/eventscheduler"
)

func main() {
	stopChannel := make(chan bool)

	// eventScheduler := eventscheduler.NewSyncEventScheduler("Sync Scheduler", stopChannel)
	eventScheduler := eventscheduler.NewConcurrentEventScheduler("Concurrent scheduler", stopChannel, 10)
	testEventScheduler(eventScheduler, 1000)
	stopChannel <- true
	fmt.Println("Exiting main")

}

func testEventScheduler(eventScheduler eventscheduler.EventSchedulerInterface, counter int) {
	var wg sync.WaitGroup
	go eventScheduler.ExecuteEventLoop()
	var eventList []eventscheduler.Event
	wg.Add(counter)
	for i := 0; i < counter; i++ {
		ev := &eventscheduler.SimpleEvent{
			Name: "Event " + strconv.FormatInt(int64(i), 10),
			Callback: func(eventName string) {
				defer wg.Done()
				fmt.Println("Starting event ", time.Now(), eventName)
				time.Sleep(time.Second * 2)
				fmt.Println("Completing event ", time.Now(), eventName)

			},
		}
		eventList = append(eventList, ev)
	}

	for _, ev := range eventList {
		eventScheduler.AddEvent(ev)
	}

	wg.Wait()
}
