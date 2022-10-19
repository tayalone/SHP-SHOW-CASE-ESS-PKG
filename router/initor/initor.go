package initor

import (
	router "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router/fiber"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router/gin"
)

/*Init Reouter Router Instant */
func Init(rType string, conf router.Config) router.Route {
	switch rType {
	case "GIN":
		return gin.NewMyRouter(conf)
	case "FIBER":
		return fiber.NewFiberRouter(conf)
	default:
		return gin.NewMyRouter(conf)
	}
}
