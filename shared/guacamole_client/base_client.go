// Package guacamole_client клиент до Apache Guacamole
package guacamole_client

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type GuacamoleClient struct {
	Config *GuacamoleClientConfig
	client *resty.Client
}

func NewGuacamoleClient(config *GuacamoleClientConfig, client *resty.Client) *GuacamoleClient {
	client.SetDebug(config.Debug)
	client.SetTimeout(time.Duration(config.MaxTimeoutSeconds) * time.Second)
	client.SetBaseURL(config.BaseUrl)
	return &GuacamoleClient{
		Config: config,
		client: client,
	}
}

func NewGuacamoleError(msg string, statusCode int) error {
	if msg == "" {
		msg = "Unknown error"
	}
	return fmt.Errorf("GuacamoleClient: %s; statusCode: %d", msg, statusCode)
}
