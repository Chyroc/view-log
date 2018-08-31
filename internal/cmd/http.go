package cmd

import (
	"net/http"

	"github.com/Chyroc/vlog/internal/common"
	"github.com/Chyroc/vlog/internal/server"

	"github.com/go-chi/chi"
	"github.com/igm/sockjs-go/sockjs"
)

func NewServer() *chi.Mux {
	app := chi.NewRouter()

	app.Get("/", IndexHandler)
	server.FileServer(app, "/", http.Dir("./internal/static/"))
	app.Handle("/ws/*", WebsocketHandler())

	return app
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	server.RenderHTML(w, r, "index", map[string]interface{}{"Server": common.Config.HTTP.Server})
}

func WebsocketHandler() http.Handler {
	return sockjs.NewHandler("/ws", sockjs.DefaultOptions, func(session sockjs.Session) {
		for {
			if msg, err := session.Recv(); err == nil {
				session.Send(msg)
				continue
			}
			break
		}
	})
}
