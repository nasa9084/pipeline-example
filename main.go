package main

import (
	"log"
	"net/http"
	"os"
)

func main() { os.Exit(exec()) }

func exec() int {
	http.HandleFunc("/", healthcheckHandler)
	http.HandleFunc("/users", userListHandler)
	http.HandleFunc("/user/1", userHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Print(err)
		return 1
	}
	return 0
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func userListHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"users": [{"user_id": 1, "username": "foo"}, {"user_id": 2, "username": "bar"}, {"user_id": 3, "username": "baz"}]}`))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"user_id": 1, "username": "foo"}`))
}
