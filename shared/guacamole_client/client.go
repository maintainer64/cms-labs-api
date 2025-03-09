package guacamole_client

type GuacamoleTokenResponse struct {
	AuthToken string `json:"authToken"`
}

func (c *GuacamoleClient) GetToken(username string, password string) (*GuacamoleTokenResponse, error) {
	const path = "/html5/api/tokens"
	var responseJson GuacamoleTokenResponse
	response, err := c.client.R().SetFormData(map[string]string{
		"username": username,
		"password": password,
	}).SetResult(&responseJson).Post(path)
	if response == nil {
		return nil, NewGuacamoleError("", 0)
	}
	if response.IsError() {
		return nil, NewGuacamoleError("guacamole authorization error", response.StatusCode())
	}
	return &responseJson, err
}
