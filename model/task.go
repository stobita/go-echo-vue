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
	insertResponseData struct {
		ID int `db:"id"`
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
	var insertResponse insertResponseData
	p := new(taskInfoJSON)
	if err := c.Bind(p); err != nil {
		fmt.Println(err)
	}
	result, db_err := sess.InsertInto(tablename_tasks).Columns("name").Values(p.Name).Exec()
	if db_err == nil {
		id, _ := result.LastInsertId()
		insertResponse.ID = int(id)
	}
	return c.JSON(http.StatusOK, insertResponse)
}
