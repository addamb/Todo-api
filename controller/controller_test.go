package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/addamb/todo-api/controller"
	"github.com/addamb/todo-api/internal/handler"
	"github.com/addamb/todo-api/mocks"
	"github.com/addamb/todo-api/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	name = "buy food"
)

func Test_CreateTodo(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := map[string]struct {
		TodoJson string

		CreateTodoItemInput string
		CreateTodoItemError error

		ExpectedStatus int
		ExpectedResult string
		ExpectError    bool
	}{
		"error bind": {
			TodoJson:       "not-json",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: "bad request",
		},
		"error validate": {
			TodoJson:       `{"id":1}`,
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: `Key: 'TodoInput.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
		},
		"error creating todo item": {
			TodoJson:            fmt.Sprintf(`{"name":"%s"}`, name),
			CreateTodoItemInput: name,
			CreateTodoItemError: errors.New("test create todo item error"),
			ExpectedStatus:      http.StatusBadRequest,
			ExpectedResult:      "test create todo item error",
		},

		"success": {
			TodoJson:            fmt.Sprintf(`{"name":"%s"}`, name),
			CreateTodoItemInput: name,
			ExpectedStatus:      http.StatusAccepted,
			ExpectedResult:      "success",
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockQueries := mocks.NewQueries(t)
			h := handler.NewHandler(mockQueries)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/create", strings.NewReader(tcase.TodoJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tcase.CreateTodoItemInput != "" {
				mockQueries.EXPECT().CreateTodoItem("buy food").Return(
					tcase.CreateTodoItemError,
				)
			}

			err := h.Controller.CreateTodo(c)
			assert.Equal(tcase.ExpectError, err != nil)
			assert.Equal(tcase.ExpectedStatus, rec.Code)
			if len(rec.Body.Bytes()) > 0 {
				assert.Equal(tcase.ExpectedResult, rec.Body.String())
			}
		})
	}
}

func Test_GetTodos(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	todoItem := models.Todo{
		ID:       1,
		Name:     name,
		Finished: false,
	}

	todos := []models.Todo{
		todoItem,
	}

	expectedJson, _ := json.Marshal(todos)

	tests := map[string]struct {
		GetTodosReturns []models.Todo
		GetTodosError   error
		ExpectedStatus  int
		ExpectedResult  string
		ExpectError     bool
	}{
		"error get todos": {
			GetTodosError:  errors.New("test get todos error"),
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: "test get todos error",
		},
		"success": {
			GetTodosReturns: todos,
			ExpectedStatus:  http.StatusOK,
			ExpectedResult:  string(expectedJson) + "\n",
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockQueries := mocks.NewQueries(t)
			h := handler.NewHandler(mockQueries)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/todos", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockQueries.EXPECT().GetTodos().Return(
				tcase.GetTodosReturns,
				tcase.GetTodosError,
			)

			err := h.Controller.GetTodos(c)
			assert.Equal(tcase.ExpectError, err != nil)
			assert.Equal(tcase.ExpectedStatus, rec.Code)
			if len(rec.Body.Bytes()) > 0 {
				assert.Equal(tcase.ExpectedResult, rec.Body.String())
			}
		})
	}
}

func Test_UpdateTodoByID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	todoItem := models.Todo{
		ID:       1,
		Name:     name,
		Finished: true,
	}

	marshaledJson, _ := json.Marshal(todoItem)

	tests := map[string]struct {
		TodoJson string

		UpdateTodoItemInput models.Todo
		UpdateTodoItemError error

		ExpectedStatus int
		ExpectedResult string
		ExpectError    bool
	}{
		"error bind": {
			TodoJson:       "not-json",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: "bad request",
		},
		"error validate": {
			TodoJson:       `{"id":1}`,
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: `Key: 'TodoInput.Name' Error:Field validation for 'Name' failed on the 'required' tag`,
		},
		"error updating todo item": {
			TodoJson:            string(marshaledJson),
			UpdateTodoItemInput: todoItem,
			UpdateTodoItemError: errors.New("test update todo item error"),
			ExpectedStatus:      http.StatusBadRequest,
			ExpectedResult:      "test update todo item error",
		},
		"success": {
			TodoJson:            string(marshaledJson),
			UpdateTodoItemInput: todoItem,
			ExpectedStatus:      http.StatusAccepted,
			ExpectedResult:      "success",
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockQueries := mocks.NewQueries(t)
			h := handler.NewHandler(mockQueries)

			//not used
			controller.NewController(mockQueries)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/api/update", strings.NewReader(tcase.TodoJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if (tcase.UpdateTodoItemInput != models.Todo{}) {
				mockQueries.EXPECT().UpdateTodoItem(
					tcase.UpdateTodoItemInput.ID,
					tcase.UpdateTodoItemInput.Name,
					tcase.UpdateTodoItemInput.Finished).Return(
					tcase.UpdateTodoItemError,
				)
			}

			err := h.Controller.UpdateTodoByID(c)
			assert.Equal(tcase.ExpectError, err != nil)
			assert.Equal(tcase.ExpectedStatus, rec.Code)
			if len(rec.Body.Bytes()) > 0 {
				assert.Equal(tcase.ExpectedResult, rec.Body.String())
			}
		})
	}
}

func Test_DeleteTodosByIDs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := map[string]struct {
		TodoJson string

		DeleteTodoItemsInput []int
		DeleteTodoItemsError error

		ExpectedStatus int
		ExpectedResult string
		ExpectError    bool
	}{
		"error bind": {
			TodoJson:       "not-json",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: "bad request",
		},
		"error 0 ids": {
			TodoJson:       `{"ids": []}`,
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: "bad request",
		},
		"error updating todo item": {
			TodoJson:             `{"ids": [1,4,7]}`,
			DeleteTodoItemsInput: []int{1, 4, 7},
			DeleteTodoItemsError: errors.New("test delete todo items error"),
			ExpectedStatus:       http.StatusBadRequest,
			ExpectedResult:       "test delete todo items error",
		},
		"success": {
			TodoJson:             `{"ids": [1,4,7]}`,
			DeleteTodoItemsInput: []int{1, 4, 7},
			ExpectedStatus:       http.StatusAccepted,
			ExpectedResult:       "success",
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockQueries := mocks.NewQueries(t)
			h := handler.NewHandler(mockQueries)

			//not used
			controller.NewController(mockQueries)

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/api/delete", strings.NewReader(tcase.TodoJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if len(tcase.DeleteTodoItemsInput) > 0 {
				mockQueries.EXPECT().DeleteTodoItems(tcase.DeleteTodoItemsInput).Return(
					tcase.DeleteTodoItemsError,
				)
			}

			err := h.Controller.DeleteTodosByIDs(c)
			assert.Equal(tcase.ExpectError, err != nil)
			assert.Equal(tcase.ExpectedStatus, rec.Code)
			if len(rec.Body.Bytes()) > 0 {
				assert.Equal(tcase.ExpectedResult, rec.Body.String())
			}
		})
	}
}
