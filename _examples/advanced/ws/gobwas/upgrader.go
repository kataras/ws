package gobwas

import (
	"net/http"

	"github.com/kataras/fastws/_examples/advanced/ws"

	gobwas "github.com/gobwas/ws"
)

func Upgrader(upgrader gobwas.HTTPUpgrader) ws.Upgrader {
	return func(w http.ResponseWriter, r *http.Request) (ws.Socket, error) {
		underline, _, _, err := upgrader.Upgrade(r, w)
		if err != nil {
			return nil, err
		}

		return newSocket(underline, false), nil
	}
}
