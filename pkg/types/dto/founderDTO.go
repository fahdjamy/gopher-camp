package dto

type FounderResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	LinkedIn    string `json:"linkedIn"`
	LastUpdated string `json:"lastUpdated"`
}
