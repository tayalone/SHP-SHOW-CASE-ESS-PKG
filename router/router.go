package router

import "net/http"

/*Context is Behavior of Route Context In Application*/
type Context interface {
	Next()
	JSON(int, interface{})
	BindURI(interface{}) error
}

/*Route is Behavior of Route Method In Application*/
type Route interface {
	Start()
	GET(path string, handlers ...func(Context))
	Group(path string, handlers ...func(Context)) RouteGouping
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

/*RouteGouping is Behavior of Route Method In Application*/
type RouteGouping interface {
	GET(path string, handlers ...func(Context))
}

/*Config is Stric of  Configuraiont Of Http Router*/
type Config struct {
	Port int
}
