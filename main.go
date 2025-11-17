package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const webDir = ".\\web\\"

func main() {
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	fmt.Println("RUN")
	err := http.ListenAndServe(":7540", r)
	if (err != nil) {
		fmt.Printf("Listen and serve server error: %s\n", err.Error())
	}
}