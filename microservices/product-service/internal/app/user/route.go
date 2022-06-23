package user

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
}
