package store

import (
	"errors"
	"sync"

	"crud-microservice/models"
)

// (In-memory "database")

type UserStore struct {
	mu    sync.RWMutex
	users map[string]models.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]models.User),
	}
}

func (s *UserStore) Create(user models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.ID] = user
}

func (s *UserStore) Get(id string) (models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, exists := s.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserStore) GetAll() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	users := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *UserStore) Update(id string, updated models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.users[id]
	if !exists {
		return errors.New("user not found")
	}
	updated.ID = id // ensure ID stays the same
	s.users[id] = updated
	return nil
}

func (s *UserStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
