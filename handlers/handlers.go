package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NoistNT/go-first-practice/api"
	"github.com/gorilla/mux"
)

// ---------- USERS ----------  //
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal users JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := api.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal user JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	posts, err := api.GetUserPosts(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal User Posts JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetUserTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todos, err := api.GetUserTodos(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal User Todos JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}


// ---------- POSTS ----------  //
func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := api.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal posts JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	id := vars["id"]

	post, err := api.GetPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal post JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetPostComments(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	id := vars["id"]

	postComments, err := api.GetPostComments(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(postComments, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal PostComments JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}


// ---------- POSTS ----------  //
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := api.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal Todos JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	id := vars["id"]

	todo, err := api.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(todo, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marhsal Todo JSON", http.StatusInternalServerError)
	} 

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}


// ---------- ALBUMS ----------  //
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := api.GetAlbums()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(albums, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal comments JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	id := vars["id"]

	album, err := api.GetAlbum(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(album, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal Album JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetAlbumPhotos(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	id := vars["id"]

	photos, err := api.GetAlbumPhotos(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(photos, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal Album Photos JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}


// ---------- PHOTOS ----------  //
func GetPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := api.GetPhotos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(photos, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal Photos JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	id := vars["id"]

	photo, err := api.GetPhoto(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(photo, "", "  ")
	if err != nil {
		http.Error(w, "Failed to Marshal Photo JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}