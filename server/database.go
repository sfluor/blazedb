package server

import (
	"fmt"
)

type database struct {
	memory map[string][]byte
}

func newDatabase() *database {
	return &database{map[string][]byte{}}
}

func (d *database) get(key string) ([]byte, error) {
	val, ok := d.memory[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}

	return val, nil
}

func (d *database) set(key string, value []byte) error {
	_, ok := d.memory[key]
	if ok {
		return fmt.Errorf("Warning %s already has a value", key)
	}
	d.memory[key] = value
	return nil
}

func (d *database) update(key string, value []byte) error {
	_, ok := d.memory[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}
	d.memory[key] = value
	return nil
}

func (d *database) delete(key string) error {
	_, ok := d.memory[key]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}

	delete(d.memory, key)

	return nil
}