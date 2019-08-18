package calendar

import (
	"github.com/golang/protobuf/ptypes"
	"testing"
	"time"
)

func TestCalendar_CRUD(t *testing.T) {
	const (
		MethodCreate = iota
		MethodDelete
		MethodUpdate
		MethodGetEvent
		MethodGetEvents
	)
	c := NewCalendar()
	date, _ := ptypes.TimestampProto(time.Date(2020,
		12, 15, 0, 30, 10, 0, time.Local))
	tests := []struct {
		name    string
		method  int
		event   *Event
		wantErr bool
		want    int32
	}{
		{
			"OK Create 1",
			MethodCreate,
			&Event{
				Name: "Event 1",
				Type: Event_EVENT,
				Date: date,
			},
			false,
			0,
		},
		{
			"OK Create 2",
			MethodCreate,
			&Event{
				Name: "Event 2",
				Type: Event_MEETING,
				Date: date,
			},
			false,
			1,
		},
		{
			"OK Create 3 with Id",
			MethodCreate,
			&Event{
				Name: "Event 3",
				Id:   0,
				Type: Event_EVENT,
				Date: date,
			},
			false,
			2,
		},
		{
			"OK Delete 3",
			MethodDelete,
			&Event{
				Id: 1,
			},
			false,
			0,
		},
		{
			"Fail Delete 3",
			MethodDelete,
			&Event{
				Id: 1,
			},
			true,
			0,
		},
		{
			"OK Update 3",
			MethodUpdate,
			&Event{
				Id:   2,
				Name: "New Event 3",
				Type: Event_EVENT,
				Date: date,
			},
			false,
			0,
		},
		{
			"OK GetEvent 3",
			MethodGetEvent,
			&Event{
				Id:   2,
				Name: "New Event 3",
			},
			false,
			0,
		},
		{
			"Fail GetEvent 4",
			MethodGetEvent,
			&Event{
				Id:   3,
				Name: "New Event 4",
			},
			true,
			-1,
		},
		{
			"OK Create 4",
			MethodCreate,
			&Event{
				Name: "Event 4",
				Type: Event_EVENT,
				Date: date,
			},
			false,
			3,
		},
		{
			"OK GetEvents",
			MethodGetEvents,
			nil,
			false,
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				got    int32
				err    error
				method string
			)

			switch tt.method {
			case MethodCreate:
				method = "CreateEvent()"
				got, err = c.CreateEvent(tt.event)
			case MethodDelete:
				method = "CreateEvent()"
				err = c.DeleteEvent(tt.event.Id)
				got = 0
			case MethodUpdate:
				method = "UpdateEvent()"
				err = c.UpdateEvent(tt.event)
				got = 0
			case MethodGetEvent:
				method = "GetEvent()"
				ev, er := c.GetEvent(tt.event.Id)
				got = -1
				if ev != nil && er == nil && ev.Name == tt.event.Name {
					got = 0
				}
				err = er
			case MethodGetEvents:
				method = "GetEvents()"
				ev, er := c.GetEvents()
				got = -1
				if ev != nil && er == nil {
					got = int32(len(ev))
				}
				err = er
			}

			t.Log(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", method, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("%s got = %v, want %v", method, got, tt.want)
			}
		})
	}
}
