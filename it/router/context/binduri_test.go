package context

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/it/router/mock"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

type CTXBindURITestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *CTXBindURITestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *CTXBindURITestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *CTXBindURITestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-ctx-binduri/1", nil)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'JSON' OK!!",
		"id":      1,
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *CTXBindURITestSuite) TestGin() {
	suite.runTest()
}

func (suite *CTXBindURITestSuite) TestFiber() {
	suite.runTest()
}

/*TestCTXBindURISuiteSuite is trigger run it test*/
func TestCTXBindURISuiteSuite(t *testing.T) {
	suite.Run(t, new(CTXBindURITestSuite))
}
