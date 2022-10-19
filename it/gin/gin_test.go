package gin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	routers "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/routers"
	RouteInitor "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/routers/initor"
)

func initRouter() {}

/*TestSuite is a test suit for Repo*/
type TestSuite struct {
	suite.Suite
	router routers.Route
}

func iSayPing(c routers.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

func iPassFromNext(c routers.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "next is working good!!",
	})
}

func myCustomMdw(c routers.Context) {
	c.Next()
}

/*SetupSuite init setup for BookRepo*/
func (suite *TestSuite) SetupSuite() {
	// suite.Router = router.SetUpRouter()

	myRouter := RouteInitor.Init("GIN", routers.Config{Port: 3000})

	myRouter.GET("/test-get", func(c routers.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'GET' OK!!",
		})
	})

	v1 := myRouter.Group("/v1")
	v1.GET("/test-group-get", func(c routers.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test  Route Grouping 'GET' OK!!",
		})
	})

	myRouter.GET("/test-ctx-json", func(c routers.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test CTX 'JSON' OK!!",
		})
	})

	myRouter.GET("/test-ctx-next", func(c routers.Context) {
		c.Next()
	}, func(c routers.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test CTX 'Next' OK!!",
		})
	})

	myRouter.GET("/test-ctx-binduri/:id", func(c routers.Context) {
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

	suite.router = myRouter
}

/*TestRouteGet Test Gin Route GET is  Working ...*/
func (suite *TestSuite) TestRouteGet() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test-get", nil)
	suite.router.ServeHTTP(w, req)

	wantMap := map[string]interface{}{
		"message": "Test Route 'GET' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(want), w.Body.String())
}

/*TestRouteGroupingGet Test Gin Route Grouping GET is  Working ...*/
func (suite *TestSuite) TestRouteGroupingGet() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/test-group-get", nil)
	suite.router.ServeHTTP(w, req)

	wantMap := map[string]interface{}{
		"message": "Test  Route Grouping 'GET' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(want), w.Body.String())
}

/*TestRouteCTXNext Test Gin CTX Next is  Working ...*/
func (suite *TestSuite) TestRouteCTXNext() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test-ctx-next", nil)
	suite.router.ServeHTTP(w, req)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'Next' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(want), w.Body.String())
}

/*TestRouteCTXJson Test Gin CTX JSON is  Working ...*/
func (suite *TestSuite) TestRouteCTXJson() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test-ctx-json", nil)
	suite.router.ServeHTTP(w, req)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'Next' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(want), w.Body.String())
}

/*TestRouteCTXBindURI Test Gin CTX BindURI is  Working ...*/
func (suite *TestSuite) TestRouteCTXBindURI() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test-ctx-binduri/1", nil)
	suite.router.ServeHTTP(w, req)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'JSON' OK!!",
		"id":      1,
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, w.Code)
	suite.JSONEq(string(want), w.Body.String())
}

/*TestGinRouteSuite is trigger run it test*/
func TestGinRouteSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
