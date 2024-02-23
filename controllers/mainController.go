package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
  return c.String(http.StatusOK, "Hello There");
}

func GetPost(c echo.Context) error {
}

func GetPosts(c echo.Context) error {

}

func CreatePost(c echo.Context) error {

}

func UpdatePost(c echo.Context) error {

}

func ArchivePost(c echo.Context) error {

}

func DeletePost(c echo.Context) error {

}