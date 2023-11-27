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
func (c *Client) GetDataConnections(filter models.Filter) (*models.GetConnectionsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-connections", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	q.Add("limit", strconv.Itoa(filter.Limit))

	req.URL.RawQuery = q.Encode()

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	connections := models.GetConnectionsResponse{}
	err = json.Unmarshal(body, &connections)
	if err != nil {
		return nil, err
	}

	return &connections, nil
}

// GetSpaceByName - Returns list of coffees (no auth required)
func (c *Client) GetDataConnection(dataConnectionId string) (*models.GetConnectionResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-connections/%s", c.HostURL, dataConnectionId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	connection := models.GetConnectionResponse{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// CreateSpace - Creates a new space
func (c *Client) CreateDataConnection(newDataConnection models.ConnectionCreate) (*models.ConnectionCreateResponse, error) {
	rb, err := json.Marshal(newDataConnection)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/data-connections", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	dataConnection := models.ConnectionCreateResponse{}
	err = json.Unmarshal(body, &dataConnection)
	if err != nil {
		return nil, err
	}

	return &dataConnection, nil
}

func (c *Client) UpdateDataConnection(dataConnectionID string, updateDataConnection models.ConnectionUpdate) error {
	rb, err := json.Marshal(updateDataConnection)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/data-connections/%s", c.HostURL, dataConnectionID), strings.NewReader(string(rb)))
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

func (c *Client) DeleteDataConnection(dataConnectionID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/data-connections/%s", c.HostURL, dataConnectionID), nil)
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
