package devicecontrol

import (
	"net"
	"sync"

	"github.com/labstack/echo"
	"github.com/nsyszr/ngvpn/pkg/manager"
)

// Handler contains all properties to serve the API
type Handler struct {
	mgr      manager.Manager
	sessions map[net.Conn]*Session
	sync.RWMutex
}

// NewHandler create a new API handler
func NewHandler(mgr manager.Manager) *Handler {
	return &Handler{
		mgr:      mgr,
		sessions: make(map[net.Conn]*Session),
	}
}

// RegisterRoutes attaches the handlers to the echo web server
func (h *Handler) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/devicecontrol")
	api.Any("/v1", h.websocketHandler())
}
