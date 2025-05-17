package routes

import (
	"testing"

	"gitlab.com/a10869/api-modules/pnetlabaddon/app/usecases"

	"github.com/stretchr/testify/assert"

	"github.com/h2non/gock"
)

func externalAttemptsMock() {
	cmsCorePingResponse := map[string]interface{}{
		"error": false,
		"result": map[string]interface{}{
			"count": 1,
		},
	}

	gock.New("https://cms-core.local").
		MatchHeader("Authorization", "Basic Y2xpZW50LWlkOnBuZXRsYWI=").
		MatchHeader("X-Client-ID", "client-id").
		Post("/api/v1/pnet-server/ping").
		Reply(200).
		JSON(cmsCorePingResponse)
}

func TestPNETServerPingSuccess(t *testing.T) {
	description := "second factor set cookies. "
	defer gock.Off()
	externalAttemptsMock()
	f := NewFiberTestHTTP()

	expectedCode := 204
	statusCode, _ := f.Request(&FiberTestHttpRequest{
		Method: "POST",
		Route:  "/pnet-lab-addon/api/v1/task/pnet-server-ping",
		Body:   FiberRequestPayload(nil),
	})
	assert.Equal(t, expectedCode, statusCode, description+"Status code")
}

func TestFilterClearOldExternalUnlFiles(t *testing.T) {
	uc := usecases.PnetServerPingUC{}
	tests := []struct {
		name            string
		osFilePaths     []string
		activeFilePaths []string
		expectedFiles   []string
	}{
		{
			name: "Deleted lists",
			osFilePaths: []string{
				"/opt/unetlab/external/1.unl",
				"/opt/unetlab/external/labs/2.unl",
				"/opt/unetlab/external/labs/3.unl",
				"/4.unl",
				"/opt/unetlab/external/5.unl",
				"/opt/unetlab/external/6.unl",
			},
			activeFilePaths: []string{
				"/external/1.unl",
				"/external/labs/2.unl",
				"/Network2025/10.unl",
				"/Network2025/1/3.unl",
				"/external/4.unl",
				"external/5.unl",
				"external/6.unl",
			},
			expectedFiles: []string{
				"/opt/unetlab/external/labs/3.unl",
				"/opt/unetlab/external/5.unl",
				"/opt/unetlab/external/6.unl",
			},
		},
		{
			name:            "Empty lists",
			osFilePaths:     []string{},
			activeFilePaths: []string{},
			expectedFiles:   []string{},
		},
	}
	for _, test := range tests {
		result := uc.FilterClearOldExternalUnlFiles(test.osFilePaths, test.activeFilePaths)
		assert.Equal(t, test.expectedFiles, result, test.name)
	}
}
