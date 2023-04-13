package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	User{"user1", "password1"},
	User{"user2", "password2"},
}

func main() {
	router := mux.NewRouter()

	// API endpoint for adding data
	router.HandleFunc("/data", addData).Methods("POST")

	// API endpoint for retrieving data
	router.HandleFunc("/data", getData).Methods("GET")

	// start server
	http.ListenAndServe(":8080", router)

	// exercises
	// 1. Написати вебсервер, який по GET отримує 2 числа і знак (*/+-), а відповідає результатом цієї операції
}

// middleware function for authentication
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, user := range users {
			if user.Username == username && user.Password == password {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}

// middleware function for authorization
func authzMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add authorization logic here
		// for example, check if the user has access to the requested resource
		// if not, return a 403 Forbidden status code

		next.ServeHTTP(w, r)
	})
}

// handler function for adding data
func addData(w http.ResponseWriter, r *http.Request) {
	// check if the user is authenticated and authorized
	authMiddleware(authzMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// add data to database or do some other processing
		fmt.Println(data)

		w.WriteHeader(http.StatusCreated)
	}))).ServeHTTP(w, r)
}

// handler function for getting data
func getData(w http.ResponseWriter, r *http.Request) {
	// check if the user is authenticated and authorized
	authMiddleware(authzMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// retrieve data from database or do some other processing
		data := []map[string]interface{}{
			{"id": 1, "name": "data1"},
			{"id": 2, "name": "data2"},
			{"id": 3, "name": "data3"},
		}

		// return data as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}))).ServeHTTP(w, r)
}
