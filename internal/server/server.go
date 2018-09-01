package server

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Chyroc/vlog/internal/common"
)

type Server struct {
	*chi.Mux
}

func New() *Server {
	app := &Server{Mux: chi.NewRouter()}

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		RenderHTML(w, r, "index", map[string]interface{}{"Server": common.Config.HTTP.Server})
	})
	app.Command("/cmd")
	app.Websocket("/ws")
	app.StaticFile("/", "./internal/static/")

	return app
}
