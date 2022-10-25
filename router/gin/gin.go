package gin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	router "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

/*MyGinContext is Overide gin contexts*/
type MyGinContext struct {
	*gin.Context
}

/*Next use in Middleware */
func (c *MyGinContext) Next() {
	c.Context.Next()
}

/*JSON return everything to json*/
func (c *MyGinContext) JSON(statuscode int, v interface{}) {
	c.Context.JSON(statuscode, v)
}

/*BindURI return error uri*/
func (c *MyGinContext) BindURI(obj interface{}) error {
	err := c.Context.ShouldBindUri(obj)
	return err
}

/*NewMyGinContext create My New Context*/
func NewMyGinContext(c *gin.Context) *MyGinContext {
	return &MyGinContext{Context: c}
}

/*NewGinHandler covert  MyGinContext -> Gin Context */
func NewGinHandler(handler func(c router.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyGinContext(c))
	}
}

// MyGinRouter is Overided Gin Engine
type MyGinRouter struct {
	*gin.Engine
	conf router.Config
}

// NewMyRouter retun my engin
func NewMyRouter(conf router.Config) *MyGinRouter {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"*",
	}
	config.AllowHeaders = []string{}
	r.Use(cors.New(config))

	return &MyGinRouter{r, conf}
}

func handlerConvertor(h []func(router.Context)) []gin.HandlerFunc {
	ginHandlers := []gin.HandlerFunc{}
	for _, handler := range h {
		ginHandlers = append(ginHandlers, NewGinHandler(handler))
	}
	return ginHandlers
}

/*Start is Command Gin Router Start */
func (r *MyGinRouter) Start() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", r.conf.Port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("listen: %s\n", err)
		} else {
			fmt.Println("Gin Running @ port", r.conf.Port)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	fmt.Println("Server exiting")
}

// GET Hadeler HTTP gin
func (r *MyGinRouter) GET(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)

	r.Engine.GET(path, ginHandlers...)
}

/*Group  Routing*/
func (r *MyGinRouter) Group(path string, handlers ...func(router.Context)) router.RouteGouping {
	ginHandlers := handlerConvertor(handlers)
	return MyGinRouterGroup{RouterGroup: r.Engine.Group(path, ginHandlers...)}
}

/*MyGinRouterGroup .... */
type MyGinRouterGroup struct {
	*gin.RouterGroup
}

/*GET .... */
func (r MyGinRouterGroup) GET(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	r.RouterGroup.GET(path, ginHandlers...)
}

func (r *MyGinRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Engine.ServeHTTP(w, req)
}

/*Testing make Gin Testing Call API and return result and statuscode*/
func (r *MyGinRouter) Testing(method string, path string, body map[string]interface{}) (int, string) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, _ := http.NewRequest(method, path, b)
	w := httptest.NewRecorder()
	r.Engine.ServeHTTP(w, req)

	return w.Code, w.Body.String()
}

/*TestServeHTTP do joblike Testing*/
func (r *MyGinRouter) TestServeHTTP(method string, path string, body map[string]interface{}) (int, string) {
	return r.Testing(method, path, body)
}
