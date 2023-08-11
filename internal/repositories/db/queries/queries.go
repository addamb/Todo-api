package queries

import (
	"database/sql"
	"time"

	"github.com/addamb/todo-api/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Queries interface {
	CreateTodoItem(name string) error
	GetTodos() ([]models.Todo, error)
	UpdateTodoItem(id int, name string, status bool) error
	DeleteTodoItems(ids []int) error
}

type SqlXQueries struct {
	DBX *sqlx.DB
}

func NewSQLXQueries(db *sqlx.DB) *SqlXQueries {
	return &SqlXQueries{
		DBX: db,
	}
}

func (s *SqlXQueries) CreateTodoItem(name string) (err error) {
	query := `INSERT INTO public.todo ("name") VALUES ($1);`
	_, err = s.DBX.Exec(query, name)
	if err != nil {
		return err
	}

	return
}

func (s *SqlXQueries) GetTodos() (todos []models.Todo, err error) {
	query := `SELECT id, name, finished, created_at, updated from todo`

	// run query
	rows, err := s.DBX.Query(query)
	if err != nil {
		return
	}

	// iterate over each row
	for rows.Next() {
		var id int
		var name string
		var finished bool
		var createdAt time.Time
		var updated sql.NullTime

		err = rows.Scan(&id, &name, &finished, createdAt, updated)

		todos = append(todos, models.Todo{
			ID:        id,
			Name:      name,
			Finished:  finished,
			CreatedAt: createdAt,
			Updated:   updated.Time,
		})
	}

	// check the error from rows
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (s *SqlXQueries) UpdateTodoItem(id int, name string, status bool) (err error) {
	query := `UPDATE public.todo
		SET "name"=$2, finished=$3, 
		updated=current_timestamp
		WHERE id= $1;`
	_, err = s.DBX.Exec(query, id, name, status)
	if err != nil {
		return err
	}

	return
}

func (s *SqlXQueries) DeleteTodoItems(ids []int) (err error) {
	query := `DELETE FROM public.todo WHERE id = ANY($1);`
	_, err = s.DBX.Exec(query, pq.Array(ids))
	if err != nil {
		return err
	}

	return
}
