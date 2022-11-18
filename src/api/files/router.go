package files

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (api *fileApi) Router() http.Handler {
	r := chi.NewRouter()
	r.Post("/upload", api.upload)
	return r
}
