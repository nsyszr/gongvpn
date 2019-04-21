package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nsyszr/ngvpn/pkg/model"
)

func (h *Handler) setGroupRoutes(api *echo.Group) {
	api.OPTIONS("/v1/groups", corsPreflightHandler)
	api.GET("/v1/groups", h.groupListHandler)
	api.POST("/v1/groups", h.groupCreateHandler)

	api.OPTIONS("/v1/groups/:id", corsPreflightHandler)
	api.GET("/v1/groups/:id", h.groupGetByIDHandler)
}

func (h *Handler) groupListHandler(c echo.Context) error {
	resourceList := []model.Group{}

	groups, err := h.mgr.Groups().FetchAll()
	if err != nil {
		return err
	}

	for _, group := range groups {
		resourceList = append(resourceList, group)
	}

	return c.JSON(http.StatusOK, resourceList)
}

func (h *Handler) groupCreateHandler(c echo.Context) error {
	group := &model.Group{}
	if err := c.Bind(group); err != nil {
		return err
	}

	if err := h.mgr.Groups().Create(group); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, group)
}

func (h *Handler) groupGetByIDHandler(c echo.Context) error {
	id := c.Param("id")

	group, err := h.mgr.Groups().FindByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, group)
}
