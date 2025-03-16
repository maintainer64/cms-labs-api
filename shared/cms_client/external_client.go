package cms_client

func (c *CMSClient) PnetServerPing(params *PNETServerPingInputDTO) (*PNETServerPingOutputDTO, error) {
	const path = "/api/v1/pnet-server/ping"
	var ssoToken PNETServerPingResponse
	response, err := c.client.R().SetBody(params).SetBasicAuth(c.Config.ClientID, c.Config.Token).SetResult(&ssoToken).Post(path)
	if response == nil {
		return nil, NewCMSError("", 0)
	}
	if response.IsError() || ssoToken.Error {
		return nil, NewCMSError(ssoToken.Msg, response.StatusCode())
	}
	return &ssoToken.Result, err
}
