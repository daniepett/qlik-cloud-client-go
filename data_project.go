package qlikcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/qlik-cloud-client-go/models"
)

// // GetSpaces - Returns list of coffees (no auth required)
// func (c *Client) GetDataConnections(filter models.Filter) (*models.DataConnections, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-connections", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	q := req.URL.Query()

// 	q.Add("limit", strconv.Itoa(filter.Limit))

// 	req.URL.RawQuery = q.Encode()

// 	fmt.Println("This is the requestss", req)
// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	connections := models.DataConnections{}
// 	err = json.Unmarshal(body, &connections)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &connections, nil
// }

// GetSpaceByName - Returns list of coffees (no auth required)
func (c *Client) GetDataProject(dataProjectId string) (*models.DataProjectResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-projects/%s", c.HostURL, dataProjectId), nil)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	project := models.DataProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

// CreateSpace - Creates a new space
func (c *Client) CreateDataProject(newDataProject models.DataProjectCreate) (*models.DataProjectResponse, error) {
	rb, err := json.Marshal(newDataProject)

	println(string(rb))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/data-projects", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	project := models.DataProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) UpdateDataProject(dataProjectID string, updateDataProject models.DataProjectUpdate) (*models.DataProjectResponse, error) {
	rb, err := json.Marshal(updateDataProject)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/data-projects/%s", c.HostURL, dataProjectID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	project := models.DataProjectResponse{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) DeleteDataProject(dataProjectID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/data-projects/%s", c.HostURL, dataProjectID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
