package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type Users struct {
	Email       string `json:"email" form:"email"`
	Nama        string `json:"nama" form:"nama"`
	NoHandphone string `json:"no_handphone" form:"no_handphone"`
	ALamat      string `json:"alamat" form:"alamat"`
	Ktp         string `json:"ktp" form:"ktp"`
}

func main() {
	route := echo.New()
	route.POST("user/create_user", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)
		contentType := c.Request().Header.Get("Content-type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
		} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
			file, err := c.FormFile("ktp")
			if err != nil {
				fmt.Println("Ktp kosong")
			} else {
				src, err := file.Open()
				if err != nil {
					return err
				}
				defer src.Close()
				dst, err := os.Create(file.Filename)
				if err != nil {
					return err
				}
				defer dst.Close()
				if _, err = io.Copy(dst, src); err != nil {
					return err
				}

				user.Ktp = file.Filename
				fmt.Println("Ada file, akan disimpan")
			}
		}
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses melakukan penambahan data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:email", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)
		user.Email = c.Param("email")
		// do something here ....
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses mengupdate data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_user/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		// do something here ....
		response := struct {
			Message string
			ID      string
		}{
			Message: "Sukses menghapus data",
			ID:      user.Email,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/search_user", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.QueryParam("keywords")
		user.Nama = "Abu Zubayr"
		user.ALamat = "Jalan Jalan"
		user.Ktp = "file.jpg"
		// do something here ....
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses melihat data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.Start(":9000")
}
