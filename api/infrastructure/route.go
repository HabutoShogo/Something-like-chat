package infrastructure

import (
	"sample/interfaces/controllers"
	"github.com/labstack/echo"
)

func Init() {
	// Echo instance
	e := echo.New()

	SampleController := controllers.NewSampleController(NewSqlHandler())

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/samples", func(c echo.Context) error { return SampleController.Index(c) })
	e.GET("/sample/:id", func(c echo.Context) error { return SampleController.Show(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

