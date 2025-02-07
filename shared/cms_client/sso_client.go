package cms_client

import (
	"fmt"

	"github.com/google/uuid"
)

func (c *CMSClient) SSOToken(grantType string, redirectUri string, code string, refreshToken string) (*SSOToken, error) {
	const url = "/api/v1/sso/token"
	var ssoToken SSOTokenResponse
	response, err := c.client.R().SetBody(map[string]string{
		"grant_type":    grantType,
		"redirect_uri":  redirectUri,
		"code":          code,
		"refresh_token": refreshToken,
	}).SetResult(&ssoToken).Post(url)
	if response == nil {
		return nil, NewCMSError("", 0)
	}
	if response.IsError() || ssoToken.Error {
		return nil, NewCMSError(ssoToken.Msg, response.StatusCode())
	}
	return &ssoToken.Result, err
}

func (c *CMSClient) SSOUserInfo(accessToken string) (*SSOTokenPublicData, error) {
	const url = "/v1/sso/userinfo"
	var ssoToken SSOTokenPublicDataResponse
	response, err := c.client.R().SetHeader(
		"Authorization",
		fmt.Sprintf("Bearer %s", accessToken),
	).SetResult(&ssoToken).Post(url)
	if response == nil {
		return nil, NewCMSError("", 0)
	}
	if response.IsError() || ssoToken.Error {
		return nil, NewCMSError(ssoToken.Msg, response.StatusCode())
	}
	return &ssoToken.Result, err
}

func (c *CMSClient) SSOAuthorizeURI(
	redirectUri string,
	scope string,
	path string,
	extra string,
) string {
	const url = "/api/v1/sso/authorize"
	state := uuid.New().String()
	if path == "" {
		path = "/"
	}
	return fmt.Sprintf(
		"%s%s?redirect_uri=%s&scope=%s&state=%s&path=%s&extra=%s",
		c.client.BaseURL,
		url,
		redirectUri,
		scope,
		state,
		path,
		extra,
	)
}
