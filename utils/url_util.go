package utils

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UrlParams(r *http.Request, path string) string {
	return chi.URLParam(r, path)
}

func UrlQuery(r *http.Request, value string) string {
	return r.URL.Query().Get(value)
}
