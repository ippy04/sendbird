package sendbird

type RequestDefaults struct {
	Auth string `json:"auth,omitempty"`
}

func (s *RequestDefaults) PopulateAuthApiToken(client *SendbirdClient) {
	s.Auth = client.ApiToken
}

type RequestDefaultsAPIV2 struct {
	ApiToken string `json:"api_token"`
}

func (s *RequestDefaultsAPIV2) PopulateApiV2Token(client *SendbirdClient) {
	s.ApiToken = client.ApiToken
}
