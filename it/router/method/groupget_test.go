package method

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/it/router/mock"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

type GroupGetTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *GroupGetTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *GroupGetTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "Fiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *GroupGetTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/v1/test-group-get", nil)

	wantMap := map[string]interface{}{
		"message": "Test  Route Grouping 'GET' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *GroupGetTestSuite) TestGin() {
	suite.runTest()
}

func (suite *GroupGetTestSuite) TestFiber() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestGroupRouteGetSuite(t *testing.T) {
	suite.Run(t, new(GroupGetTestSuite))
}
