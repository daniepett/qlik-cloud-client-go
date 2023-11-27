package qlikcloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/qlik-cloud-client-go/models"
)

// GetSpaces - Returns list of coffees (no auth required)
func (c *Client) GetConnectionString(src string, props models.GetConnectionString, pass models.GetConnectionString) (*models.GetConnectionStringResponse, error) {
	rb, err := json.Marshal(props)

	rp, err := json.Marshal(pass.PropertiesList)

	payload := &models.GetConnectionRequest{Array: []string{string(rb), "", string(rp)}}

	enc, err := json.Marshal(payload.Array)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/dcaas/command?name=getConnectionString&dataSourceId=%s", c.HostURL, src), strings.NewReader(string(enc)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	res := models.GetConnectionStringResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
