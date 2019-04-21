package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nsyszr/ngvpn/pkg/manager"
)

// Handler contains all properties to serve the API
type Handler struct {
	mgr manager.Manager
}

// NewHandler create a new API handler
func NewHandler(mgr manager.Manager) *Handler {
	return &Handler{mgr: mgr}
}

// RegisterRoutes attaches the handlers to the echo web server
func (h *Handler) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	h.setGroupRoutes(api)
}

func corsPreflightHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
