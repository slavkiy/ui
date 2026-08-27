package reactive

import "sync"

// Subscription represents a cancellable registration.
type Subscription struct {
	unsubscribe func()
	once        sync.Once
}

// Unsubscribe removes the subscription. It is safe to call more than once.
func (s *Subscription) Unsubscribe() {
	if s == nil || s.unsubscribe == nil {
		return
	}
	s.once.Do(s.unsubscribe)
}
