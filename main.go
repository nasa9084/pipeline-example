package main

import (
	"log"
	"net/http"
	"os"
)

func main() { os.Exit(exec()) }

func exec() int {
	http.HandleFunc("/", healthcheckHandler)
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
