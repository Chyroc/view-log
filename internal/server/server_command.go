package server

import (
	"io"
	"net/http"

	"github.com/Chyroc/vlog/internal/command"
)

func (r *Server) Command(pre string) {
	var commands []string
	for k := range command.Commands {
		commands = append(commands, k)
		r.Post(pre+"/"+k, func(w http.ResponseWriter, r *http.Request) {
			var args []string
			if err := BindJSON(r, &args); err != nil {
				if err != io.EOF {
					RenderError(w, r, err)
					return
				}
			}

			out, err := command.RunWithSlice(append([]string{k}, args...), false)
			if err != nil {
				RenderError(w, r, err)
				return
			}

			RenderSuccess(w, r, out)
		})
	}

	r.Get(pre, func(w http.ResponseWriter, r *http.Request) {
		RenderSuccess(w, r, commands)
	})
}
