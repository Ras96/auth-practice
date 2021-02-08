package model

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        uuid.UUID `json:"id,omitempty"  db:"ID"`
	Title     string    `json:"title,omitempty"  db:"Title"`
	CreatedAt time.Time `json:"createdAt,omitempty"  db:"CreatedAt"`
}

func GetTodoHandler(c echo.Context) error {
	todos := []Todo{}
	err := db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}
	if len(todos) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, todos)
}

func PostTodoHandler(c echo.Context) error {
	todo := Todo{}
	c.Bind(&todo)

	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM todos WHERE title = ?", todo.Title)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}

	if count > 0 {
		return c.String(http.StatusConflict, "同名のタスクが既に追加されています")
	}

	_, err = db.Exec("INSERT INTO todos (Title) VALUES (?)", todo.Title)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %v", err))
	}
	return c.NoContent(http.StatusCreated)
}
