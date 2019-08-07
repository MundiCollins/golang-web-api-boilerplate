package handlers

import (
	"database/sql"
	"github.com/MundiCollins/golang-web-api-boilerplate/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetPosts(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPosts(db))
	}
}