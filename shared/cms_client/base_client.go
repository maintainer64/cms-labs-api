// Package cms_client клиент до CMS Labs Core
package cms_client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

type CMSClient struct {
	Config  *CMSClientConfig
	BaseURL *url.URL
	client  *resty.Client
}

func NewCMSClient(config *CMSClientConfig) *CMSClient {
	client := resty.New()
	client.SetDebug(config.Debug)
	client.SetTimeout(time.Duration(config.MaxTimeoutSeconds) * time.Second)
	client.SetBaseURL(config.BaseUrl)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("X-Client-ID", config.ClientID)
	uri, _ := url.Parse(config.BaseUrl)
	return &CMSClient{
		Config:  config,
		BaseURL: uri,
		client:  client,
	}
}

func NewCMSError(msg string, statusCode int) error {
	if msg == "" {
		msg = "Unknown error"
	}
	return fmt.Errorf("CMSError: %s; statusCode: %d", msg, statusCode)
}
