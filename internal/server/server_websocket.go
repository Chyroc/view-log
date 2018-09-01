package server

import (
	"github.com/igm/sockjs-go/sockjs"
)

func (r *Server) Websocket(pre string) {
	r.Handle(pre+"/*", sockjs.NewHandler(pre, sockjs.DefaultOptions, func(session sockjs.Session) {
		for {
			if msg, err := session.Recv(); err == nil {
				session.Send(msg)
				continue
			}
			break
		}
	}))
}
