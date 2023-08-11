package handler

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRoutes(api *echo.Group) {
	api.POST("/create", h.Controller.CreateTodo)
	api.GET("/todos", h.Controller.GetTodos)
	api.PUT("/update", h.Controller.UpdateTodoByID)
	api.DELETE("/delete", h.Controller.DeleteTodosByIDs)
}
