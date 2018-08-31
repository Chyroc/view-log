package cmd

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Chyroc/vlog/internal/common"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/igm/sockjs-go/sockjs"
)

func NewServer() *chi.Mux {
	app := chi.NewRouter()

	app.Get("/", IndexHandler)
	FileServer(app, "/", http.Dir("./internal/static/"))
	app.Handle("/ws/*", WebsocketHandler())

	return app
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./internal/static/index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	html, err := ParseTmpl(string(b), map[string]interface{}{"Server": common.Config.HTTP.Server})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	render.HTML(w, r, string(html))
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

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

// ParseTmpl ...
func ParseTmpl(tmpl string, data interface{}) ([]byte, error) {
	parsedTmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := parsedTmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
