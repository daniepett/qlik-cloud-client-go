package qlikcloud

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daniepett/qlik-cloud-client-go/models"
)

// GetDataGateway - Returns list of coffees (no auth required)
func (c *Client) GetDataGateway(dataGatewayID string) (*models.DataGateway, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/replicate-agents/%s", c.HostURL, dataGatewayID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	dataGateway := models.DataGateway{}
	err = json.Unmarshal(body, &dataGateway)
	if err != nil {
		return nil, err
	}

	return &dataGateway, nil
}

// GetDataGatewayByName - Returns list of coffees (no auth required)
// func (c *Client) GetDataGatewayByName(dataGatewayName string) (*models.DataGateway, error) {
// 	req, err := http.NewRequest("GET", c.HostURL, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	qp := `filter=name co "` + dataGatewayName + `"`

// 	fmt.Println("This is the requestss", req)
// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dataGateways := models.DataGateways{}
// 	err = json.Unmarshal(body, &dataGateways)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dataGateways.DataGateways[0], nil
// }
