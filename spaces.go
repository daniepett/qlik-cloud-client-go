package qlikcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/daniepett/qlik-cloud-client-go/models"
)

// GetSpaces - Returns list of coffees (no auth required)
func (c *Client) GetSpaces(filter models.Filter) (*models.Spaces, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/spaces", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	q.Add("limit", strconv.Itoa(filter.Limit))
	q.Add("name", filter.Name)

	req.URL.RawQuery = q.Encode()

	fmt.Println("This is the requestss", req)
	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	spaces := models.Spaces{}
	err = json.Unmarshal(body, &spaces)
	if err != nil {
		return nil, err
	}

	return &spaces, nil
}

// GetSpaceByName - Returns list of coffees (no auth required)
func (c *Client) GetSpace(spaceId string) (*models.Space, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/spaces/%s", c.HostURL, spaceId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	space := models.Space{}
	err = json.Unmarshal(body, &space)
	if err != nil {
		return nil, err
	}

	return &space, nil
}

// CreateSpace - Creates a new space
func (c *Client) CreateSpace(newSpace models.CreateSpace) (*models.Space, error) {
	rb, err := json.Marshal(newSpace)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/spaces", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	space := models.Space{}
	err = json.Unmarshal(body, &space)
	if err != nil {
		return nil, err
	}

	return &space, nil
}

func (c *Client) UpdateSpace(spaceID string, updatedSpace models.UpdateSpace) (*models.Space, error) {
	rb, err := json.Marshal(updatedSpace)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/spaces/%s", c.HostURL, spaceID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	space := models.Space{}
	err = json.Unmarshal(body, &space)
	if err != nil {
		return nil, err
	}

	return &space, nil
}

func (c *Client) DeleteSpace(spaceID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/spaces/%s", c.HostURL, spaceID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return err
	}

	if string(body) != "" {
		return errors.New(string(body))
	}

	return nil
}
