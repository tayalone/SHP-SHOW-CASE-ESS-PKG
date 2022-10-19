package main

import (
	"fmt"
	"net/http"

	routers "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/routers"
	RouteInitor "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/routers/initor"
)

type tmpRoute struct {
	routers.Route
}

var mtr tmpRoute

func newRoute() routers.Route {
	myRouter := RouteInitor.Init("GET", routers.Config{Port: 3000})
	myRouter.GET("/test-get", func(c routers.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'GET' OK!!",
		})
	})
	mtr.Route = myRouter
	return mtr
}

func main() {
	fmt.Println("xxx")
	myRouter := newRoute()

	myRouter.Start()
}
