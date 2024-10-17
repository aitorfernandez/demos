package button

import "testing"

func TestAddEventListener(t *testing.T) {
	chFoo := make(chan string, 1)
	chBar := make(chan string, 1)

	e := "test"

	btn := New()
	btn.AddEventListener(e, chFoo)
	btn.AddEventListener(e, chBar)

	if _, ok := btn.EventListeners[e]; !ok {
		t.Errorf("want an event got %v", ok)
	}

	if len(btn.EventListeners[e]) != 2 {
		t.Errorf("want 2 got %v", len(btn.EventListeners[e]))
	}
}

func TestRemoveEventListener(t *testing.T) {
	ch := make(chan string, 1)

	e := "test"

	btn := New()
	btn.AddEventListener(e, ch)
	btn.RemoveEventListener(e, ch)

	if len(btn.EventListeners[e]) != 0 {
		t.Errorf("want 0 got %v", len(btn.EventListeners[e]))
	}
}
