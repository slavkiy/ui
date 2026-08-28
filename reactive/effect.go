package reactive

// Effect registers a callback that receives each new signal value.
func (s *Signal[T]) Effect(fn func(T)) *Subscription {
	if fn == nil {
		return &Subscription{}
	}
	s.mu.Lock()
	id := s.nextID
	s.nextID++
	s.effects[id] = fn
	s.mu.Unlock()
	return &Subscription{unsubscribe: func() {
		s.mu.Lock()
		delete(s.effects, id)
		s.mu.Unlock()
	}}
}

// SubscribeEffect registers a callback that runs when the signal changes.
func (s *Signal[T]) SubscribeEffect(fn func()) *Subscription {
	if fn == nil {
		return &Subscription{}
	}
	return s.Effect(func(T) { fn() })
}
