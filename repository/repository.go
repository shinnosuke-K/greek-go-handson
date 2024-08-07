package repository

import "errors"

type Repository interface {
	Store(key string, value string) error
	Get(key string) (string, error)
}

type InMemoryRepository struct {
	data map[string]string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: map[string]string{},
	}
}

func (r *InMemoryRepository) Store(key string, value string) error {
	r.data[key] = value
	return nil
}

func (r *InMemoryRepository) Get(key string) (string, error) {
	value, ok := r.data[key]
	if !ok {
		return "", errors.New("not found")
	}
	return value, nil
}
