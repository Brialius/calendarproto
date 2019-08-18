//go:generate protoc -I=$IMPORT_PATH -I=. --go_out=. calendar.proto

package calendar

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"strings"
	"sync/atomic"
)

type Calendar struct {
	*EventsMap
	id int32
}

func (c *Calendar) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("Calendar %p\n", c))
	for _, event := range c.Events {
		date, _ := ptypes.Timestamp(event.Date)
		created, _ := ptypes.Timestamp(event.CreatedAt)
		updated, _ := ptypes.Timestamp(event.LastUpdated)
		s.WriteString(fmt.Sprintf("id: %d, name: %s, event type: %s, event date: %s, created: %s, updated: %s\n",
			event.Id, event.GetName(), event.GetType(), date.Local(), created.Local(), updated.Local()))
	}
	return s.String()
}

func (c *Calendar) CreateEvent(e *Event) (int32, error) {
	e.Id = atomic.LoadInt32(&c.id)
	now := ptypes.TimestampNow()
	e.CreatedAt = now
	e.LastUpdated = now
	if len(e.GetName()) == 0 {
		return 0, fmt.Errorf("event name is not set")
	}
	if _, ok := c.Events[c.id]; !ok {
		c.Events[c.id] = e
	} else {
		return 0, fmt.Errorf("event with id %d already exist", c.id)
	}
	atomic.AddInt32(&c.id, 1)
	return e.Id, nil
}

func (c *Calendar) UpdateEvent(e *Event) error {
	id := e.Id
	e.LastUpdated = ptypes.TimestampNow()
	if _, ok := c.Events[id]; !ok {
		return fmt.Errorf("event with id %d doesn't exist", c.id)
	}
	if e.Date != nil {
		c.Events[id].Date = e.Date
	}
	if len(e.Name) > 0 {
		c.Events[id].Name = e.Name
	}
	c.Events[id].Type = e.Type

	return nil
}

func (c *Calendar) GetEvent(id int32) (*Event, error) {
	if _, ok := c.Events[id]; !ok {
		return nil, fmt.Errorf("event with id %d doesn't exist", c.id)
	}
	return c.Events[id], nil
}

func (c *Calendar) GetEvents() ([]*Event, error) {
	events := make([]*Event, 0, len(c.Events))
	for _, event := range c.Events {
		events = append(events, event)
	}
	return events, nil
}

func (c *Calendar) DeleteEvent(id int32) error {
	if _, ok := c.Events[id]; !ok {
		return fmt.Errorf("event with id %d doesn't exist", id)
	}
	delete(c.Events, id)
	return nil
}

func NewCalendar() *Calendar {
	return &Calendar{
		EventsMap: &EventsMap{
			Events: map[int32]*Event{},
		},
		id: 0,
	}
}
