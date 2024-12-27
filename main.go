package main

import (
	"errors"
	"fmt"
)

type employe struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employe) error
	get(id int) (employe, error)
	delete(id int) error
}

// zanosym structuru v pamyat'
type memoryStorage struct {
	data map[int]employe
}

// realisuem structuru v pamyat'
func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employe),
	}
}

// zanosym infu v pamyat'
func (s *memoryStorage) insert(e employe) error {
	s.data[e.id] = e

	return nil
}

// poluchaem infu from memory, if id in memori doesn't exist - get error
func (s *memoryStorage) get(id int) (employe, error) {
	e, exists := s.data[id]
	if !exists {
		return employe{}, errors.New("employe doesn't exist")
	}

	return e, nil
}

// delete employe by id from memory
func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func main() {
	var s storage

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newMemoryStorage()

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)
}
