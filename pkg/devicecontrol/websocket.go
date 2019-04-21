package devicecontrol

import (
	"io"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo"
	"github.com/nsyszr/ngvpn/pkg/devicecontrol/proto"
)

func (h *Handler) websocketHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
		if err != nil {
			return err
		}

		// Add session to session map
		h.Lock()
		defer h.Unlock()
		h.sessions[conn] = NewSession()

		// Start listening the websocket connection
		go listen(conn)

		return nil
	}
}

func listen(conn net.Conn) error {
	r := wsutil.NewReader(conn, ws.StateServerSide)
	w := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	decoder := proto.NewDecoder(r)
	// encoder := protocol.NewEncoder(w)

	// Start listening for next frames
	for {
		hdr, err := r.NextFrame()
		if err != nil {
			return err
		}

		if hdr.OpCode == ws.OpClose {
			return io.EOF
		}

		var req []interface{}
		if err := decoder.Decode(&req); err != nil {
			return err
		}

		// res := handleMessage(req)
		// if err := encoder.Encode(&res); err != nil {
		//     return err
		//}

		if err := w.Flush(); err != nil {
			return err
		}
	}
}
