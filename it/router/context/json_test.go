package context

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/it/router/mock"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

type CTXJsonTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *CTXJsonTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *CTXJsonTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "Fiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *CTXJsonTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-ctx-json", nil)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'JSON' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *CTXJsonTestSuite) TestGin() {
	suite.runTest()
}

func (suite *CTXJsonTestSuite) TestFiber() {
	suite.runTest()
}

/*TestCTXNextSuite is trigger run it test*/
func TestCTXJsonSuite(t *testing.T) {
	suite.Run(t, new(CTXJsonTestSuite))
}
