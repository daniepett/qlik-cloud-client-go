package models

type Spaces struct {
	Spaces []Space `json:"data"`
	Meta   Meta    `json:"meta"`
	Links  Links   `json:"links"`
}

type Space struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	OwnerID     string `json:"ownerId"`
}

type CreateSpace struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type UpdateSpace struct {
	Name        string `json:"name"`
	OwnerID     string `json:"ownerId"`
	Description string `json:"description"`
}
