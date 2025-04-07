package store

import (
	"desafio-CRUD-MEMORIA_Rocketseat/model"
	"errors"
	"log"
	"sync"

	"github.com/google/uuid"
)


type MemoryStore struct {
	mu    sync.RWMutex
	users map[uuid.UUID]model.User
}

// NewMemoryStore cria uma nova instância do armazenamento em memória.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[uuid.UUID]model.User),
	}
}

//Insert adiciona um novo usuário ao armazenamento, gerando um novo UUID.
// utilizando o sync.Mutex para garantir a integridade dos dados (evitando data-races!)
func (s *MemoryStore) Insert(user model.User) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = uuid.New()
	s.users[user.ID] = user

	return user, nil	
}

//FindAll retorta todos os usuarios em memória
// utilizando o sync.Mutex para garantir a integridade dos dados (evitando data-races!)
func (s *MemoryStore) FindByID(id uuid.UUID) (model.User, bool){
	log.Printf("Procurando usuário com ID: %s", id)
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[id]
	log.Printf("Pegando o valor do id: %+v", user)
	return user, ok
}

//FindByID busca um usuário pelo seu ID.
// utilizando o sync.Mutex para garantir a integridade dos dados (evitando data-races!)
func (s* MemoryStore) FindAll() ([]model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]model.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}

//Delete remove um usuário do armazenamento com base no UUID.	
// utilizando o sync.Mutex para garantir a integridade dos dados (evitando data-races!)
func (s *MemoryStore) Delete(id uuid.UUID) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return model.User{}, nil
	}
	delete(s.users, id)
	return user, nil
}

//Update modifica os dados de um usuário existente com base no UUID.
// utilizando o sync.Mutex para garantir a integridade dos dados (evitando data-races!)
func (s *MemoryStore) Update(id uuid.UUID, update model.User) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}

	user.FirstName = update.FirstName
	user.LastName = update.LastName
	user.Biography = update.Biography
	s.users[id] =user
	return user, nil
}