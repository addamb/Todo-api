package controller

import (
	"net/http"

	"github.com/addamb/todo-api/internal/repositories/db/queries"
	"github.com/addamb/todo-api/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Queries queries.Queries
}

func NewController(db queries.Queries) *Controller {
	return &Controller{
		Queries: db,
	}
}

func (c *Controller) CreateTodo(e echo.Context) (err error) {
	input := new(models.TodoInput)

	if err := e.Bind(input); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	err = validator.New().Struct(input)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	err = c.Queries.CreateTodoItem(input.Name)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.String(http.StatusAccepted, "success")
}

func (c *Controller) GetTodos(e echo.Context) (err error) {

	//Get all todos
	todos, err := c.Queries.GetTodos()
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, todos)
}

func (c *Controller) UpdateTodoByID(e echo.Context) (err error) {
	input := new(models.TodoInput)

	if err := e.Bind(input); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	err = validator.New().Struct(input)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	err = c.Queries.UpdateTodoItem(input.ID, input.Name, input.Finished)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.String(http.StatusAccepted, "success")
}

func (c *Controller) DeleteTodosByIDs(e echo.Context) (err error) {
	input := new(models.DeleteInput)

	if err := e.Bind(input); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	if len(input.IDs) == 0 {
		return e.String(http.StatusBadRequest, "bad request")
	}

	err = c.Queries.DeleteTodoItems(input.IDs)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.String(http.StatusAccepted, "success")
}
