package main

import (

	"github.com/labstack/echo"
    "bookstore/db"
    "bookstore/model"
    "bookstore/router"
)

func main() {
	sql := &db.Sql{
		Host:     "34.69.17.6",
		Port:     5432,
		UserName: "choll102",
		Password: "123456",
		DbName:   "back-end-muasachonline",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, model.Response{
			StatusCode: 200,
			Message:    "Home Page",
		})
	})

	router.UserRouter(e, sql)
	router.CateRouter(e, sql)
	router.ProductRouter(e, sql)
	router.OrderRouter(e, sql)

	e.Logger.Fatal(e.Start(":3000"))
}
