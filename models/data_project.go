package models

type DataProjectConfiguration struct {
	ID                string `json:"id"`
	LakehouseType     string `json:"lakehouseType"`
	Type              string `json:"type"`
	StorageConnection string `json:"storageConnection"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	BatchMode         bool   `json:"batchMode"`
}
type DataProjectCreate struct {
	SpaceID string                   `json:"spaceId"`
	Data    DataProjectConfiguration `json:"data"`
}

type DataProjectResponse struct {
	Key         string `json:"key"`
	Info        Info   `json:"info"`
	DataProject struct {
		LakehouseType            string `json:"lakehouseType"`
		Type                     string `json:"type"`
		StorageConnection        string `json:"storageConnection"`
		LandingStorageConnection string `json:"landingStorageConnection"`
		QvdStorageLocationType   string `json:"qvdStorageLocationType"`
		BatchMode                bool   `json:"batchMode"`
		ID                       string `json:"id"`
		Name                     string `json:"name"`
		Description              string `json:"description"`
		Version                  string `json:"version"`
		Tags                     string `json:"tags"`
	} `json:"dataProject"`
}

type DataProjectUpdate struct {
	SpaceID string `json:"spaceId"`
	Data    DataProjectConfiguration
}
