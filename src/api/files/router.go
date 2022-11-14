package files

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (f *fileApi) Router() http.Handler {
	r := chi.NewRouter()
	r.Post("/upload", f.upload)
	return r
}
