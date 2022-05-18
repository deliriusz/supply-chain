package model

type ProductMetadataDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Image       string   `json:"image"`
	ExternalUri string   `json:"external_uri,omitempty"`
	Attributes  []string `json:"attributes,omitempty"`
}
