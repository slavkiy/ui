package reactive

// Persist creates a signal initialized from storage and saves subsequent values.
func Persist[T any](storage Storage, key string, defaultValue T) *Signal[T] {
	value := defaultValue
	if err := storage.Get(key, &value); err != nil {
		value = defaultValue
	}
	signal := NewSignal("", value)
	signal.Subscribe(func(value T) { _ = storage.Set(key, value) })
	return signal
}
