package model

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

var (
	conn, _ = dbr.Open("mysql", "goechovue:goechovue@tcp(localhost:3306)/goechovue", nil)
)

type (
	userInfo struct {
		Email string `db:"email"`
	}
	userInfoJSON struct {
		Email string `db:"email"`
	}
	responseData struct {
		Users []userInfo
	}
)

var (
	tablename = "users"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func InsertUser(c echo.Context) error {
	sess := conn.NewSession(nil)
	p := new(userInfoJSON)
	if err := c.Bind(p); err != nil {
		return err
	}
	sess.InsertInto(tablename).Columns("email").Values(p.Email).Exec()

	return c.NoContent(http.StatusOK)
}
