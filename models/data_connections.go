package models

type GetConnectionsResponse struct {
	Connections []GetConnectionResponse `json:"data"`
	Meta        Meta                    `json:"meta"`
	Links       Links                   `json:"links"`
}

type GetConnectionResponse struct {
	ID                  string `json:"qID"`
	Qri                 string `json:"qri"`
	User                string `json:"user"`
	Links               Links  `json:"links"`
	Name                string `json:"qName"`
	Type                string `json:"qType"`
	SpaceID             string `json:"space"`
	Tenant              string `json:"tenant"`
	Created             string `json:"created"`
	Updated             string `json:"updated"`
	DataSourceID        string `json:"datasourceID"`
	Architecture        int    `json:"qArchitecture"`
	ReferenceKey        string `json:"qReferenceKey"`
	CredentialsID       string `json:"qCredentialsID"`
	EngineObjectID      string `json:"qEngineObjectID"`
	CredentialsName     string `json:"qCredentialsName"`
	ConnectStatement    string `json:"qConnectStatement"`
	ConnectionSecret    string `json:"qConnectionSecret"`
	SeparateCredentials bool   `json:"qSeparateCredentials"`
}

type ConnectionCreate struct {
	ID                  string `json:"qID"`
	Owner               string `json:"owner"`
	Name                string `json:"qName"`
	Type                string `json:"qType"`
	SpaceID             string `json:"space"`
	LogOn               int64  `json:"qLogOn"`
	Password            string `json:"qPassword"`
	Username            string `json:"qUsername"`
	DataSourceID        string `json:"datasourceID"`
	QriInRequest        string `json:"qriInRequest"`
	Architecture        int    `json:"qArchitecture"`
	CredentialsID       string `json:"qCredentialsID"`
	EngineID            string `json:"qEngineObjectID"`
	CredentialsName     string `json:"qCredentialsName"`
	ConnectStatement    string `json:"qConnectStatement"`
	ConnectionSecret    string `json:"qConnectionSecret"`
	SeparateCredentials bool   `json:"qSeparateCredentials"`
}

type ConnectionCreateResponse struct {
	ID                  string   `json:"qID"`
	User                string   `json:"user"`
	Links               Links    `json:"links"`
	Name                string   `json:"qName"`
	Type                string   `json:"qType"`
	SpaceID             string   `json:"space"`
	LogOn               int64    `json:"qLogOn"`
	Created             string   `json:"created"`
	Updated             string   `json:"updated"`
	Privileges          []string `json:"privileges"`
	Architecture        int      `json:"qArchitecture"`
	ReferenceKey        string   `json:"qReferenceKey"`
	CredentialsID       string   `json:"qCredentialsID"`
	EngineID            string   `json:"qEngineObjectID"`
	CredentialsName     string   `json:"qCredentialsName"`
	ConnectStatement    string   `json:"qConnectStatement"`
	SeparateCredentials bool     `json:"qSeparateCredentials"`
}

type ConnectionUpdate struct {
	ID                  string `json:"qID"`
	Name                string `json:"qName"`
	Type                string `json:"qType"`
	SpaceID             string `json:"space"`
	LogOn               int64  `json:"qLogOn"`
	Password            string `json:"qPassword"`
	Username            string `json:"qUsername"`
	DataSourceID        string `json:"datasourceID"`
	Architecture        int    `json:"qArchitecture"`
	CredentialsID       string `json:"qCredentialsID"`
	EngineID            string `json:"qEngineObjectID"`
	CredentialsName     string `json:"qCredentialsName"`
	ConnectStatement    string `json:"qConnectStatement"`
	SeparateCredentials bool   `json:"qSeparateCredentials"`
	ConnectionSecret    string `json:"qConnectionSecret"`
}
