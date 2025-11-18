package api

import "github.com/go-chi/chi/v5"

const dateFormat = "20060102"

func Init(r chi.Router) {
	r.Get("/nextdate", nextDateHandler)
}