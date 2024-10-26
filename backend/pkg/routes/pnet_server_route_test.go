package routes

import (
	"io"
	"net/http/httptest"
	"testing"

	"gitlab.com/a10869/api-modules/backend/app/models"

	"github.com/stretchr/testify/assert"
)

func TestV1PNETServerGet(t *testing.T) {
	app, db := FiberAppTest()
	entity := models.PNETServer{}
	db.Create(entity)

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description  string
		route        string // input route
		method       string // input method
		body         io.Reader
		expectedBody string
		expectedCode int
	}{
		{
			description: "get book",
			route:       "/api/v1/pnet-server/get",
			body: FiberRequestPayload(map[string]any{
				"id": entity.ID,
			}),
			expectedBody: "",
			expectedCode: 200,
		},
		{
			description: "not found book",
			route:       "/api/v1/book",
			body: FiberRequestPayload(map[string]any{
				"id": 999999,
			}),
			expectedBody: "",
			expectedCode: 200,
		},
	}

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, _ := app.Test(req, -1) // the -1 disables request latency

		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
		assert.Equal(t, test.body, resp.Body, test.description)
	}
}
