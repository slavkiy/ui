package reactive

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// Storage defines operations for loading, saving, and deleting keyed values.
type Storage interface {
	// Get loads a value by key into value.
	Get(key string, value any) error
	// Set stores a value under key.
	Set(key string, value any) error
	// Delete removes the value stored under key.
	Delete(key string) error
}

// FileStorage stores JSON-encoded values in a file.
type FileStorage struct {
	mu   sync.Mutex
	path string
	data map[string]json.RawMessage
}

// NewFileStorage creates a file-backed storage and loads existing data.
func NewFileStorage(path string) *FileStorage {
	s := &FileStorage{path: path, data: make(map[string]json.RawMessage)}
	_ = s.load()
	return s
}

func (s *FileStorage) load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &s.data)
}

func (s *FileStorage) save() error {
	data, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0644)
}

// Get loads a JSON-encoded value by key into value.
func (s *FileStorage) Get(key string, value any) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, ok := s.data[key]
	if !ok {
		return os.ErrNotExist
	}
	return json.Unmarshal(data, value)
}

// Set JSON-encodes and stores a value under key.
func (s *FileStorage) Set(key string, value any) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	s.data[key] = data
	return s.save()
}

// Delete removes a value by key and saves the updated file.
func (s *FileStorage) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return s.save()
}
