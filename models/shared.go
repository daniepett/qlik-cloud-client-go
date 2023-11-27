package models

type Meta struct {
	Count int `json:"count"`
}

type Links struct {
	Self struct {
		Href string `json:"href"`
	} `json:"self"`
}

type Filter struct {
	Name  string `json:"name"`
	Limit int    `json:"limit"`
}

type Info struct {
	SpaceID string `json:"spaceId"`
	OwnerID string `json:"ownerId"`
	Type    string `json:"type"`
	Audit   struct {
		Created struct {
			Ts string `json:"ts"`
			By string `json:"by"`
		} `json:"created"`
		Modified struct {
			Ts string `json:"ts"`
			By string `json:"by"`
		} `json:"modified"`
	} `json:"audit"`
}
