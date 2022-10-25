package mock

import (
	"fmt"
	"net/http"

	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
	RouteInitor "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router/initor"
)

// MakeRoute create router for integration test
func MakeRoute(routeType string) router.Route {
	fmt.Println("routeType", routeType)
	myRouter := RouteInitor.Init(routeType, router.Config{Port: 3000})

	myRouter.GET("/test-get", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'GET' OK!!",
		})
	})

	v1 := myRouter.Group("/v1")
	v1.GET("/test-group-get", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test  Route Grouping 'GET' OK!!",
		})
	})

	myRouter.GET("/test-ctx-json", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test CTX 'JSON' OK!!",
		})
	})

	myRouter.GET("/test-ctx-next", func(c router.Context) {
		c.Next()
	}, func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test CTX 'Next' OK!!",
		})
	})

	myRouter.GET("/test-ctx-binduri/:id", func(c router.Context) {
		type getIDUri struct {
			ID uint `uri:"id" binding:"required"`
		}

		var gi getIDUri
		if err := c.BindURI(&gi); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test CTX 'JSON' OK!!",
			"id":      gi.ID,
		})
	})

	return myRouter
}
