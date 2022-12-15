package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jutionck/code-run-2022/config"
	"github.com/jutionck/code-run-2022/handler"
	"net/http"
	"os"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	h := handler.New(db)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/posts", h.BlogPostList)
	r.Get("/posts/{id}", h.BlogPostGet)
	r.Post("/posts", h.BlogPostCreate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server started at port " + port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
