package eventscheduler

type SimpleEvent struct {
	Name     string
	Callback func()
	id       Id
	isActive bool
}

func (e *SimpleEvent) Compare(ev Event) bool {
	return e.id == ev.EvenId() && e.Name == ev.EventName() && e.isActive == ev.IsEventActive()
}

func (e *SimpleEvent) CompareWithId(id Id) bool {
	return e.id == id
}

func (e *SimpleEvent) EvenId() Id {
	return e.id
}

func (e *SimpleEvent) SetEventId(id Id) {
	e.id = id
}

func (e *SimpleEvent) EventName() string {
	return e.Name
}

func (e *SimpleEvent) SetEventName(s string) {
	e.Name = s
}

func (e *SimpleEvent) IsEventActive() bool {
	return e.isActive
}

func (e *SimpleEvent) MarkEventActive(state bool) {
	e.isActive = state
}

func (e *SimpleEvent) Execute() {
	e.Callback()
}

var _ Event = &SimpleEvent{}
