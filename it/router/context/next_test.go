package context

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/it/router/mock"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/router"
)

type CTXNextTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *CTXNextTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *CTXNextTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *CTXNextTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-ctx-next", nil)

	wantMap := map[string]interface{}{
		"message": "Test CTX 'Next' OK!!",
	}
	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *CTXNextTestSuite) TestGin() {
	suite.runTest()
}

func (suite *CTXNextTestSuite) TestFiber() {
	suite.runTest()
}

/*TestCTXNextSuiteSuite is trigger run it test*/
func TestCTXNextSuiteSuite(t *testing.T) {
	suite.Run(t, new(CTXNextTestSuite))
}
