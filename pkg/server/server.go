package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kdd4/go_final_project/pkg/api"
)

func Run(webDir string) error {
	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	r.Route("/api", api.Init)

	fmt.Println("RUN")

	err := http.ListenAndServe(":7540", r)

	return err
}