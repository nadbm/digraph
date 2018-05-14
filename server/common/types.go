package common

type Organization struct {
	ExternalId string `json:"externalId,omitempty"`
	Uid        string `json:"uid,omitempty"`
	Name       string `json:"name,omitempty"`
}

type User struct {
	ExternalId string `json:"externalId,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
}

type Topic struct {
	Uid         string `json:"uid,omitempty"`
	ExternalId  string `json:"externalId,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Link struct {
	Uid   string `json:"uid,omitempty"`
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
}

type LinkConnection *[]Link
