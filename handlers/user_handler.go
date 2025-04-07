package handlers

import (
	"desafio-CRUD-MEMORIA_Rocketseat/model"
	"desafio-CRUD-MEMORIA_Rocketseat/store"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

//UserHandler é responsável por lidar com requisições relacionadas aos usuários
type UserHandler struct {
	store *store.MemoryStore
}

//NewUserHandler retorna uma nova instância de UserHandler
func NewUserHandler(store *store.MemoryStore) *UserHandler {
	return &UserHandler{store: store}
}

//ValidaUserInput valida os campos obrigatórios de um usuário
func validateUserInput(user model.User) error {
	if len(user.FirstName) <= 2 || len(user.FirstName) > 20 {
		return errors.New("first name must be between 2 and 20 characters")
	}
	if len(user.LastName) <= 2 || len(user.LastName) > 20 {
		return errors.New("last name must be between 2 and 20 characters")
	}
	if len(user.Biography) <= 2 || len(user.Biography) > 450 {
		return errors.New("biography must be between 2 and 450 characters")
	}	
	return nil	
}

//CreateUser cria um novo usuário
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	var user model.User
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validateUserInput(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := h.store.Insert(user)
	if err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}

//GetUserByID busca um usuário pelo ID informado
func (h* UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
	log.Printf("ID recebido: '%s'", idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Erro ao fazer parse do UUID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, found := h.store.FindByID(id)
	if !found {
		log.Printf("Usuário com ID %s não encontrado", id)
		http.Error(w, "Usuário nao encontrado", http.StatusNotFound	)
		return
	}
	log.Printf("Usuário encontrado: %+v", user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//GetAllUsers busca todos os usuários
func (h* UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request){
	users, err := h.store.FindAll();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

//DeleteUser deleta um usuário pelo ID informado
func (h* UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
	log.Printf("ID recebido: '%s'", idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Erro ao fazer parse do UUID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.store.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Usuário deletado: %+v", user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//UpdateUser atualiza um usuário pelo ID informado
func (h* UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
	log.Printf("ID recebido: '%s'", idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Erro ao fazer parse do UUID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)	
		return	
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validateUserInput(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return	
	}

	user, err = h.store.Update(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Usuário atualizado: %+v", user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}