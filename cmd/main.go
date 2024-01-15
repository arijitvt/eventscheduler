package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/arijit/eventscheduler/eventscheduler"
)

func main() {
	stopChannel := make(chan bool)

	// eventScheduler := eventscheduler.NewSyncEventScheduler("Sync Scheduler", stopChannel)
	eventScheduler := eventscheduler.NewConcurrentEventScheduler("Concurrent scheduler", stopChannel, 5)
	testEventScheduler(eventScheduler)
	stopChannel <- true

}

func testEventScheduler(eventScheduler eventscheduler.EventSchedulerInterface) {
	var wg sync.WaitGroup
	go eventScheduler.ExecuteEventLoop()

	wg.Add(2)
	eventList := []eventscheduler.Event{
		&eventscheduler.SimpleEvent{Name: "First", Callback: func() {
			defer wg.Done()
			fmt.Println("Executing first event ")
			time.Sleep(time.Second * 2)
			fmt.Println("First Event is done ")
		}},
		&eventscheduler.SimpleEvent{Name: "Second", Callback: func() {
			defer wg.Done()
			fmt.Println("Executing second event ")
			time.Sleep(time.Second * 1)
			fmt.Println("Second Event is done ")

		}},
	}

	for _, ev := range eventList {
		eventScheduler.AddEvent(ev)
	}

	wg.Wait()
}
