package models

type GetConnectionString struct {
	PropertiesList []ConnectionProperties `json:"propertiesList"`
}
type ConnectionProperties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetConnectionStringResponse struct {
	SessionID                   string `json:"sessionId"`
	Success                     bool   `json:"success"`
	ConnectionSecretString      string `json:"connectionSecretString"`
	ConnectionString            string `json:"connectionString"`
	CredentialsConnectionString string `json:"credentialsConnectionString"`
	DataSourceID                string `json:"dataSourceId"`
	UserID                      string `json:"userId"`
}

type GetConnectionRequest struct {
	Array []string
}
