package service

import (
	"errors"
	"sync"

	"github.com/muhammedarifp/grpc-example/pb"
)

var (
	ErrAldredyExist = errors.New("data is aldredy exist")
)

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryStore struct {
	mutex sync.RWMutex
	DB    map[string]*pb.Laptop
}

func NewInmemoryStore() *InMemoryStore {
	return &InMemoryStore{
		DB: make(map[string]*pb.Laptop),
	}
}

func (i *InMemoryStore) Save(laptop *pb.Laptop) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.DB[laptop.Id] != nil {
		return ErrAldredyExist
	}

	i.DB[laptop.Id] = laptop
	return nil
}
