package handler

import (
	"github.com/addamb/todo-api/controller"
	"github.com/addamb/todo-api/internal/repositories/db/queries"
)

type Handler struct {
	Controller *controller.Controller
}

func NewHandler(queries queries.Queries) *Handler {
	return &Handler{
		Controller: &controller.Controller{
			Queries: queries,
		},
	}
}
