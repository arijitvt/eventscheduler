package eventscheduler

type Id int64

type Event interface {
	EvenId() Id
	SetEventId(Id)
	EventName() string
	SetEventName(string)
	IsEventActive() bool
	MarkEventActive(bool)
	Compare(Event) bool
	CompareWithId(Id) bool
	Execute()
}
