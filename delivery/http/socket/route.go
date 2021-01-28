package socket

import (
	"github.com/labstack/echo/v4"
	"gorm-v2/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase}

	socket := group.Group("/ws")
	socket.GET("", r.socket)
}
