package model

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type (
	taskInfo struct {
		Name string `db:"name"`
	}
	taskInfoJSON struct {
		Name string `db:"name"`
	}
	responseTaskData struct {
		Tasks []taskInfo
	}
)

var (
	tablename_tasks = "tasks"
)

func InsertTask(c echo.Context) error {
	sess := conn.NewSession(nil)
	p := new(taskInfoJSON)
	if err := c.Bind(p); err != nil {
		fmt.Println(err)
	}
	sess.InsertInto(tablename_tasks).Columns("name").Values(p.Name).Exec()
	return c.NoContent(http.StatusOK)
}
