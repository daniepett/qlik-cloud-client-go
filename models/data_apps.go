package models

type DataAppsResponse struct {
	DataApps []DataApp `json:"dataApps"`
}

type DataAppResponse struct {
	Key     string  `json:"key"`
	Info    Info    `json:"info"`
	DataApp DataApp `json:"dataApp"`
}

type DataApp struct {
	Type        string `json:"type"`
	SubType     string `json:"subType"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Tags        string `json:"tags"`
}

type DataAppCreate struct {
	Data DataAppCreateData `json:"data"`
}

type DataAppCreateData struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type DataAppUpdate struct {
	Data DataAppUpdateData `json:"data"`
}

type DataAppUpdateData struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
