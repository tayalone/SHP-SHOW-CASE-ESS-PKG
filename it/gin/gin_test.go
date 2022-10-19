package gin

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	router "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
	RouteInitor "github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router/initor"
)

/*TestSuite is a test suit for Repo*/
type TestSuite struct {
	suite.Suite
	router router.Route
}

func iSayPing(c router.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

func iPassFromNext(c router.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "next is working good!!",
	})
}

func myCustomMdw(c router.Context) {
	c.Next()
}

/*SetupSuite init setup for BookRepo*/
func (suite *TestSuite) SetupSuite() {
	// suite.Router = router.SetUpRouter()

	myRouter := RouteInitor.Init("GIN", router.Config{Port: 3000})

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

	suite.router = myRouter
}

/*TestRouteGet Test Gin Route GET is  Working ...*/
func (suite *TestSuite) TestRouteGet() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-get", nil)

	wantMap := map[string]interface{}{
		"message": "Test Route 'GET' OK!!",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

/*TestRouteGroupingGet Test Gin Route Grouping GET is  Working ...*/
func (suite *TestSuite) TestRouteGroupingGet() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/v1/test-group-get", nil)

	wantMap := map[string]interface{}{
		"message": "Test  Route Grouping 'GET' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

/*TestRouteCTXNext Test Gin CTX Next is  Working ...*/
func (suite *TestSuite) TestRouteCTXNext() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-ctx-next", nil)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'Next' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

/*TestRouteCTXJson Test Gin CTX JSON is  Working ...*/
func (suite *TestSuite) TestRouteCTXJson() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-ctx-json", nil)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'JSON' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

/*TestRouteCTXBindURI Test Gin CTX BindURI is  Working ...*/
func (suite *TestSuite) TestRouteCTXBindURI() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-ctx-binduri/1", nil)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'JSON' OK!!",
		"id":      1,
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

/*TestGinRouteSuite is trigger run it test*/
func TestGinRouteSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
