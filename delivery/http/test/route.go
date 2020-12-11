package test

import (
	"gorm-v2/usecase"

	"github.com/labstack/echo/v4"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase}

	test := group.Group("/test")
	test.GET("", r.getAll)
	test.GET("/:id", r.getTestByID)
}
