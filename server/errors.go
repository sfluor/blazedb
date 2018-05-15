package server

import "fmt"

type KeyNotFoundError struct {
	key string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("Error: key %s not found", e.key)
}

type AlreadySetError struct {
	key   string
	value []byte
}

func (e AlreadySetError) Error() string {
	return fmt.Sprintf("Error: key %s already has value %s", e.key, e.value)
}
