package qlikcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetToken - Get a new token for user
func (c *Client) GetToken() (*AuthResponse, error) {
	if c.Auth.ClientID == "" || c.Auth.ClientSecret == "" {
		return nil, fmt.Errorf("Missing ClientID and ClientSecret")
	}

	if c.Auth.GrantType != "client_credentials" {
		return nil, fmt.Errorf("Grant type not supported")
	}

	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", c.HostURL), strings.NewReader(string(rb)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
