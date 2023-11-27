package models

type SourceEntities struct {
	Entities []SourceEntity `json:"entities"`
}

type SourceEntity struct {
	ID        string `json:"entityID"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Schema    string `json:"schema"`
	Database  string `json:"database"`
	ProjectID string `json:"projectId"`
	DataAppID string `json:"dataAppId"`
}

type GetSourceEntities struct {
	SourceSelection GetSourceEntitiesSourceSelection `json:"sourceSelection"`
}

type GetSourceEntitiesSourceSelection struct {
	DataEntitiesSelection GetSourceEntitiesDataEntitiesSelection `json:"dataEntitiesSelection"`
}

type GetSourceEntitiesDataEntitiesSelection struct {
	SourceConnectionID string                             `json:"sourceConnectionId"`
	IncludePatterns    []GetSourceEntitiesIncludePatterns `json:"includePatterns"`
}

type GetSourceEntitiesIncludePatterns struct {
	ProjectID     string `json:"projectId"`
	Database      string `json:"database"`
	TablePattern  string `json:"tablePattern"`
	SchemaPattern string `json:"schemaPattern"`
	EntityType    string `json:"entityType"`
}

type SourceSelectionPut struct {
	Data SourceSelectionPutData `json:"data"`
}

type SourceSelectionPutData struct {
	DataEntitiesSelection SourceSelectionDataEntitiesSelection `json:"dataEntitiesSelection"`
}

type SourceSelectionResponse struct {
	Key             string                                    `json:"key"`
	SourceSelection SourceSelectionPutResponseSourceSelection `json:"sourceSelection"`
	Info            Info                                      `json:"info"`
}

type SourceSelectionPutResponseSourceSelection struct {
	DataEntitiesSelection SourceSelectionDataEntitiesSelection `json:"dataEntitiesSelection"`
}

type SourceSelectionDataEntitiesSelection struct {
	SourceConnectionID string         `json:"sourceConnectionId"`
	DataEntities       []SourceEntity `json:"dataEntities"`
}
