package fiber

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	routers "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/routers"
)

/*MyFiberContext is Overide fiber contexts*/
type MyFiberContext struct {
	*fiber.Ctx
}

/*Next use in Middleware */
func (c *MyFiberContext) Next() {
	c.Ctx.Next()
}

/*JSON use in Middleware */
func (c *MyFiberContext) JSON(statuscode int, v interface{}) {
	c.Ctx.Status(statuscode).JSON(v)
}

/*BindURI return everything to json*/
func (c *MyFiberContext) BindURI(obj interface{}) error {
	// c.Context.JSON(statuscode, v)
	err := c.Ctx.ParamsParser(obj)
	return err
}

/*NewMyFiberContext create My New Context*/
func NewMyFiberContext(ctx *fiber.Ctx) *MyFiberContext {
	return &MyFiberContext{Ctx: ctx}
}

/*MyFiberRouter defibne Fiber */
type MyFiberRouter struct {
	*fiber.App
	conf routers.Config
}

/*NewFiberRouter defibne Fiber Router */
func NewFiberRouter(conf routers.Config) *MyFiberRouter {
	r := fiber.New()
	return &MyFiberRouter{r, conf}
}

func handlerConvertor(h []func(routers.Context)) []func(*fiber.Ctx) error {
	fiberHandlers := []func(*fiber.Ctx) error{}
	for _, handler := range h {
		fiberHandlers = append(fiberHandlers, func(c *fiber.Ctx) error {
			handler(NewMyFiberContext(c))
			return nil
		})
	}
	return fiberHandlers
}

/*Start is Command Fiber Router Start */
func (r *MyFiberRouter) Start() {
	r.Listen(fmt.Sprintf(fmt.Sprintf(":%d", r.conf.Port)))
}

/*GET Hadeler HTTP gin */
func (r *MyFiberRouter) GET(path string, handlers ...func(routers.Context)) {
	fiberHandlers := handlerConvertor(handlers)
	r.App.Get(path, fiberHandlers...)
}

/*Group is Group Routing For Fiber */
func (r *MyFiberRouter) Group(path string, handlers ...func(routers.Context)) routers.RouteGouping {
	fiberHandlers := handlerConvertor(handlers)
	return MyFiberRouterGroup{Router: r.App.Group(path, fiberHandlers...)}
}

/*MyFiberRouterGroup .... */
type MyFiberRouterGroup struct {
	fiber.Router
}

/*GET Hadeler HTTP gin */
func (r MyFiberRouterGroup) GET(path string, handlers ...func(routers.Context)) {
	fiberHandlers := handlerConvertor(handlers)
	r.Router.Get(path, fiberHandlers...)
}

func (r *MyFiberRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/*do not thing*/
}
