package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (r *Server) API(pre string) {
	app := chi.NewRouter()
	app.Get("/ls", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Handle(pre, app)
}
