package di

import (
	resty "github.com/go-resty/resty/v2"
)

var (
	NewRestyClient = func() *resty.Client {
		return resty.New()
	}
)
