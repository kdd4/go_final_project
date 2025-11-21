package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/kdd4/go_final_project/pkg/api"
)

func Run(webDir string) error {
	port := os.Getenv("TODO_PORT")
	
	if port == "" {
		os.Setenv("TODO_PORT", "7540")
		port = os.Getenv("TODO_PORT")
	}

	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	r.Route("/api", api.Init)
 
	fmt.Printf("Run server on port %s\n", port)

	address := fmt.Sprintf(":%s", port)

	err := http.ListenAndServe(address, r)

	return err
}