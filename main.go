package main

import (
	"log"
	"net/http"

	"github.com/NoistNT/go-first-practice/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Create new router
	r := mux.NewRouter()

	// Register routes handlers
	// Users
	r.HandleFunc("/users", handlers.GetUsers)
	r.HandleFunc("/users/{id}", handlers.GetUser)
	r.HandleFunc("/users/{id}/posts", handlers.GetUserPosts)
	r.HandleFunc("/users/{id}/todos", handlers.GetUserTodos)

	// Posts
	r.HandleFunc("/posts", handlers.GetPosts)
	r.HandleFunc("/posts/{id}", handlers.GetPost)
	r.HandleFunc("/posts/{id}/comments", handlers.GetPostComments)

	// TODOS
	r.HandleFunc("/todos", handlers.GetTodos)
	r.HandleFunc("/todos/{id}", handlers.GetTodo)

	// Albums
	r.HandleFunc("/albums", handlers.GetAlbums)
	r.HandleFunc("/albums/{id}", handlers.GetAlbum)
	r.HandleFunc("/albums/{id}/photos", handlers.GetAlbumPhotos)

	// Photos
	r.HandleFunc("/photos", handlers.GetPhotos)
	r.HandleFunc("/photos/{id}", handlers.GetPhoto)

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
