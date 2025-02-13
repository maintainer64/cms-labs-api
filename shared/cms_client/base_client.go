// Package cms_client клиент до CMS Labs Core
package cms_client

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type CMSClient struct {
	debug   bool
	ClintID string
	BaseURL string
	client  *resty.Client
}

func NewCMSClient(config *CMSClientConfig) *CMSClient {
	client := resty.New()
	client.SetDebug(config.Debug)
	client.SetTimeout(time.Duration(config.MaxTimeoutSeconds) * time.Second)
	client.SetBaseURL(config.BaseUrl)
	client.SetHeader("Accept", "application/json")
	client.SetBasicAuth(config.ClientID, config.Token)
	client.SetHeader("X-Client-ID", config.ClientID)
	return &CMSClient{
		debug:   config.Debug,
		ClintID: config.ClientID,
		BaseURL: config.BaseUrl,
		client:  client,
	}
}

func NewCMSError(msg string, statusCode int) error {
	if msg == "" {
		msg = "Unknown error"
	}
	return fmt.Errorf("CMSError: %s; statusCode: %d", msg, statusCode)
}
