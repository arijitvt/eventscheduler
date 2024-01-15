package eventscheduler

type EventSchedulerInterface interface {
	AddEvent(Event) bool
	DeleteEvent(Event) bool
	DeleteEventById(Id) bool
	ExecuteEventLoop()
}
