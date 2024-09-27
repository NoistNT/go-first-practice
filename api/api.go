package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int    		 `json:"id"`
	Name     string 		 `json:"name"`
	Username string 		 `json:"username"`
	Email    string 		 `json:"email"`
	Address  struct {
		Street  string 		 `json:"street"`
		Suite   string 		 `json:"suite"`
		City    string 		 `json:"city"`
		Zipcode string 		 `json:"zipcode"`
		Geo     struct {
			Lat string 			 `json:"lat"`
			Lng string 			 `json:"lng"`
		} 								 `json:"geo"`
	} 								 	 `json:"address"`
	Phone    string 		 `json:"phone"`
	Website  string 		 `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	}
}

type Todo struct {
	UserID   			int    	`json:"userId"`
	Id       			int    	`json:"id"`
	Title    			string 	`json:"title"`
	Completed     bool 		`json:"completed"`
}

type Post struct {
	UserID   int    		 `json:"userId"`
	Id       int    		 `json:"id"`
	Title    string 		 `json:"title"`
	Body     string 		 `json:"body"`
}

type PostComment struct {
	PostID		int 			 `json:"postId"`
	Id 				int 			 `json:"id"`
	Name			string 		 `json:"name"`
	Email			string 		 `json:"email"`
	Body			string 		 `json:"body"`
}

type Album struct {
	UserID 		int				 `json:"userId"`
	Id 				int				 `json:"id"`
	Title 		string		 `json:"title"`
}

type Photo struct {
	AlbumID   		int    		 `json:"albumId"`
	Id       			int    		 `json:"id"`
	Title    			string 		 `json:"title"`
	URL     			string 		 `json:"url"`
	ThumbnailUrl  string		 `json:"thumbnailUrl"`
}


// ---------- USERS ----------  //
func GetUsers() ([]User, error) {
	var users []User
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return users, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return users, fmt.Errorf("failed to fetch users")
	}

	err = json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetUser(id string) (User, error) {
	var user User

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return user, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return user, fmt.Errorf("failed to fetch user")
	}

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserPosts(id string) ([]Post, error) {
	var posts []Post

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s/posts", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return posts, err
	}

	if response.StatusCode != http.StatusOK {
		return posts, fmt.Errorf("failed to fetch user posts")
	}

	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func GetUserTodos(id string) ([]Todo, error) {
	var todos []Todo

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s/todos", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return todos, err
	}

	if response.StatusCode != http.StatusOK {
		return todos, fmt.Errorf("failed to fetch user todos")
	}

	err = json.NewDecoder(response.Body).Decode(&todos)
	if err != nil {
		return todos, err
	}

	return todos, err
}


// ---------- POSTS ----------  //
func GetPosts() ([]Post, error) {
	var posts []Post

	reponse, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return posts, err
	}
	defer reponse.Body.Close()

	if reponse.StatusCode != http.StatusOK {
		return posts, fmt.Errorf("failed to fetch posts")
	}

	err = json.NewDecoder(reponse.Body).Decode(&posts)
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func GetPost(id string) (Post, error) {
	var post Post

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return post, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return post, fmt.Errorf("failed to fetch post")
	}

	err = json.NewDecoder(response.Body).Decode(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func GetPostComments(postId string) ([]PostComment, error) {
	var postComments []PostComment

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s/comments", postId)
	
	response, err := http.Get(apiURL)
	if err != nil {
		return postComments, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return postComments, fmt.Errorf("failed to fetch post comments")
	}

	err = json.NewDecoder(response.Body).Decode(&postComments)
	if err != nil {
		return postComments, err
	}

	return postComments, nil
}


// ---------- TODOS ----------  //
func GetTodos() ([]Todo, error) {
	var todos []Todo

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return todos, err
	}

	if response.StatusCode != http.StatusOK {
		return todos, fmt.Errorf("failed to fetch todos")
	}

	err = json.NewDecoder(response.Body).Decode(&todos)
	if err != nil {
		return todos, err
	}

	return todos, err
}

func GetTodo(id string) (Todo, error) {
	var todo Todo

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%s", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return todo, err
	}

	if response.StatusCode != http.StatusOK {
		return todo, fmt.Errorf("failed to fetch todo")
	}

	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		return todo, err
	}

	return todo, err
}


// ---------- ALBUMS ----------  //
func GetAlbums() ([]Album, error) {
	var albums []Album

	response, err := http.Get("https://jsonplaceholder.typicode.com/albums")
	if err != nil {
		return albums, err
	}

	if response.StatusCode != http.StatusOK {
		return albums, fmt.Errorf("failed to fetch comments")
	}

	err = json.NewDecoder(response.Body).Decode(&albums)
	if err != nil {
		return albums, err
	}

	return albums, nil
}

func GetAlbum(id string) (Album, error) {
	var album Album

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/albums/%s", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return album, err
	}

	if response.StatusCode != http.StatusOK {
		return album, fmt.Errorf("failed to fetch album")
	}

	err = json.NewDecoder(response.Body).Decode(&album)
	if err != nil {
		return album, err
	}

	return album, nil
}


func GetAlbumPhotos(id string) ([]Photo, error) {
	var photos []Photo

	apiURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/albums/%s/photos", id)

	response, err := http.Get(apiURL)
	if err != nil {
		return photos, err
	}

	if response.StatusCode != http.StatusOK {
		return photos, fmt.Errorf("failed to fetch album photos")
	}

	err = json.NewDecoder(response.Body).Decode(&photos)
	if err != nil {
		return photos, err
	}

	return photos, err
}


// ---------- PHOTOS ----------  //
func GetPhotos() ([]Photo, error) {
	var photos []Photo

	response, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		return photos, err
	}

	if response.StatusCode != http.StatusOK {
		return photos, fmt.Errorf("failed to fetch photos")
	}

	err = json.NewDecoder(response.Body).Decode(&photos)
	if err != nil {
		return photos, err
	}

	return photos, err
}

func GetPhoto(id string) (Photo, error) {
	var photo Photo

	apiUrl := fmt.Sprintf("https://jsonplaceholder.typicode.com/photos/%s", id)

	response, err := http.Get(apiUrl)
	if err != nil {
		return photo, err
	}

	if response.StatusCode != http.StatusOK {
		return photo, fmt.Errorf("failed to fetch photo")
	}

	err = json.NewDecoder(response.Body).Decode(&photo)
	if err != nil {
		return photo, err
	}

	return photo, err
}
