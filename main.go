package main

import (
	"desafio-CRUD-MEMORIA_Rocketseat/handlers"
	"desafio-CRUD-MEMORIA_Rocketseat/store"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	menStore := store.NewMemoryStore()
	handler := handlers.NewUserHandler(menStore)

	r.Route("/api/users", func(r chi.Router){
		r.Use(jsonMiddleware)
		r.Post("/", handler.CreateUser)
		r.Get("/{id}", handler.GetUserByID)
		r.Get("/", handler.GetAllUsers)
		r.Delete("/{id}", handler.DeleteUser)
		r.Put("/{id}", handler.UpdateUser)
	})

	log.Println("Server is running on port :8200")
	http.ListenAndServe(":8200", r)
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}