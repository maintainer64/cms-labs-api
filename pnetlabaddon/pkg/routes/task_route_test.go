package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h2non/gock"
)

func externalAttemptsMock() {
	// Mock ListAttempts - возвращаем пустой список попыток
	listAttemptsResponse := map[string]interface{}{
		"jsonrpc": "2.0",
		"result": map[string]interface{}{
			"model": []interface{}{},
		},
	}

	gock.New("https://cms-core.local").
		Post("/api/v1/rpc").
		Reply(200).
		JSON(listAttemptsResponse)
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
