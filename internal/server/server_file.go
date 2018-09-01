package server

import (
	"net/http"
	"strings"
)

func (r *Server) StaticFile(path, static string) {
	root := http.Dir(static)

	if strings.ContainsAny(path, "{}*") {
		panic("StaticFile does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Mux.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Mux.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
