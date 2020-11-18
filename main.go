package main

import (
	"sample/route"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := route.Init()
	router.Start(":3000")

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 		return c.String(fasthttp.StatusOK, "aa!")
	// })
	// e.Logger.Fatal(e.Start(":3000"))
}