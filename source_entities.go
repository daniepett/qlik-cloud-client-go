package qlikcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/qlik-cloud-client-go/models"
)

// GetSpaces - Returns list of coffees (no auth required)
func (c *Client) GetSourceEntities(project string, app string, selection models.GetSourceEntities) (*models.SourceEntities, error) {
	rb, err := json.Marshal(selection)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps/%s/source-entities", c.HostURL, project, app), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ent := models.SourceEntities{}
	err = json.Unmarshal(body, &ent)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (c *Client) GetSourceSelection(project string, app string) (*models.SourceSelectionResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps/%s/source-selection", c.HostURL, project, app), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ent := models.SourceSelectionResponse{}
	err = json.Unmarshal(body, &ent)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (c *Client) PutSourceSelection(project string, app string, selection models.SourceSelectionPut) (*models.SourceSelectionResponse, error) {
	rb, err := json.Marshal(selection)

	println(string(rb))

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/data-projects/%s/data-apps/%s/source-selection", c.HostURL, project, app), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ent := models.SourceSelectionResponse{}
	err = json.Unmarshal(body, &ent)
	if err != nil {
		return nil, err
	}

	return &ent, nil
}
