package qlikcloud

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default QlikCloud URL

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// AuthResponse -
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresAt   string `json:"expires_at"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewClient -
func NewClient(tenantId, region, clientId, clientSecret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    fmt.Sprintf("https://%s.%s.qlikcloud.com", *tenantId, *region),
	}

	// If username or password not provided, return empty client
	if clientId == nil || clientSecret == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		ClientID:     *clientId,
		ClientSecret: *clientSecret,
	}

	c.Auth.GrantType = "client_credentials"

	ar, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	c.Token = ar.AccessToken

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300

	if !statusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
