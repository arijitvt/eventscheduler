package eventscheduler

import (
	"io"
	"net/http"
)

type HttpEvent struct {
	Name     string
	Url      string
	Payload  string
	Response string
	id       Id
	isActive bool
	client   *http.Client
}

func (e *HttpEvent) Compare(ev Event) bool {
	return e.id == ev.EvenId() && e.Name == ev.EventName() && e.isActive == ev.IsEventActive()
}

func (e *HttpEvent) CompareWithId(id Id) bool {
	return e.id == id
}

func (e *HttpEvent) EvenId() Id {
	return e.id
}

func (e *HttpEvent) SetEventId(id Id) {
	e.id = id
}

func (e *HttpEvent) EventName() string {
	return e.Name
}

func (e *HttpEvent) SetEventName(s string) {
	e.Name = s
}

func (e *HttpEvent) IsEventActive() bool {
	return e.isActive
}

func (e *HttpEvent) MarkEventActive(state bool) {
	e.isActive = state
}

func (e *HttpEvent) initClient() {

}

func (e *HttpEvent) Execute() {
	if e.client == nil {
		e.initClient()
	}
	resp, err := http.Get(e.Url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	e.Response = string(respBytes)
}

var _ Event = &HttpEvent{}
