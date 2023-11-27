package qlikcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/qlik-cloud-client-go/models"
)

// GetSpaces - Returns list of coffees (no auth required)
func (c *Client) GetDataApps(project string) (*models.DataAppsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps", c.HostURL, project), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	apps := models.DataAppsResponse{}
	err = json.Unmarshal(body, &apps)
	if err != nil {
		return nil, err
	}

	return &apps, nil
}

// GetSpaceByName - Returns list of coffees (no auth required)
func (c *Client) GetDataApp(project string, app string) (*models.DataAppResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps/%s", c.HostURL, project, app), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	a := models.DataAppResponse{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// CreateSpace - Creates a new space
func (c *Client) CreateDataApp(project string, app models.DataAppCreate) (*models.DataAppResponse, error) {
	rb, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps", c.HostURL, project), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	a := models.DataAppResponse{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (c *Client) UpdateDataApp(project string, app models.DataAppUpdate) (*models.DataAppResponse, error) {
	rb, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps", c.HostURL, project), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	a := models.DataAppResponse{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (c *Client) DeleteDataApp(project string, app string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps/%s", c.HostURL, project, app), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
