package models

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Finished  bool      `json:"finished"`
	CreatedAt time.Time `json:"created_at"`
	Updated   time.Time `json:"updated"`
}

type TodoInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Finished bool   `json:"finished"`
}

type DeleteInput struct {
	IDs []int `json:"ids"`
}
