package models

type DataGateways struct {
	DataGateways []DataGateway `json:"agents"`
}

type DataGateway struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"agentType"`
	Description     string `json:"description"`
	OwnerID         string `json:"ownerId"`
	SpaceID         string `json:"spaceId"`
	State           string `json:"state"`
	GatewayID       string `json:"gatewayId"`
	Version         string `json:"version"`
	VersionMetadata struct {
		ReplicateVersion   string `json:"replicateVersion"`
		QemVersion         string `json:"qemVersion"`
		DataGatewayVersion string `json:"dataGatewayVersion"`
		ManagerVersion     string `json:"managerVersion"`
	} `json:"agentVersionMetadata"`
	Enabled bool `json:"enabled"`
}

type DataGatewayCreate struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type DataGatewayUpdate struct {
	Name        string `json:"name"`
	OwnerID     string `json:"ownerId"`
	Description string `json:"description"`
}
