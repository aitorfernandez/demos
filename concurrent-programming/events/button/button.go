package button

type Button struct {
	// Slice of channels.
	EventListeners map[string][]chan string
}

func New() *Button {
	return &Button{
		EventListeners: make(map[string][]chan string),
	}
}

// AddEventListener add channel to listeners map.
func (b *Button) AddEventListener(e string, ch chan string) {
	// A registered event can have multiple channels.
	if _, ok := b.EventListeners[e]; ok {
		b.EventListeners[e] = append(b.EventListeners[e], ch)
	} else {
		b.EventListeners[e] = []chan string{ch}
	}
}

func (b *Button) RemoveEventListener(e string, ch chan string) {
	remove := func(i int) {
		b.EventListeners[e] = append(b.EventListeners[e][:i], b.EventListeners[e][i+1:]...)
	}

	if _, ok := b.EventListeners[e]; ok {
		for i, _ := range b.EventListeners[e] {
			if b.EventListeners[e][i] == ch {
				remove(i)
				break
			}
		}
	}
}

func (b *Button) DispatchEvent(e string, res string) {
	if _, ok := b.EventListeners[e]; ok {
		for _, ch := range b.EventListeners[e] {
			go func(ch chan string) {
				ch <- res
			}(ch)
		}
	}
}
