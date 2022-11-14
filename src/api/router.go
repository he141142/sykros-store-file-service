package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (handler *ApiHandler) BaseRouter(r *chi.Mux) *ApiHandler {
	// ping
	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	})

	r.Route("/api/v1", func(r chi.Router) {
		// -- routes --
		// User route
		if handler.FileApi != nil {
			r.Mount("/file", handler.FileApi.Router())
		} else {
			fmt.Println("FilesAPI is not registered therefore skipped")
		}

	})
	return handler
}
