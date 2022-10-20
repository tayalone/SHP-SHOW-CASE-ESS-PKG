package method

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/it/router/mock"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

type GetTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *GetTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *GetTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "Fiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *GetTestSuite) testGet() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-get", nil)

	wantMap := map[string]interface{}{
		"message": "Test Route 'GET' OK!!",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *GetTestSuite) TestGin() {
	suite.testGet()
}

func (suite *GetTestSuite) TestFiber() {
	suite.testGet()
}

/*TestGinRouteSuite is trigger run it test*/
func TestRouteGetSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}
