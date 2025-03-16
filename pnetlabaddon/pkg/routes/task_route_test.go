package routes

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
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
	statusCode, _ := f.Request(
		"POST",
		"/pnet-lab-addon/api/v1/task/pnet-server-ping",
		FiberRequestPayload(nil),
		"",
	)
	assert.Equal(t, expectedCode, statusCode, description+"Status code")
}
