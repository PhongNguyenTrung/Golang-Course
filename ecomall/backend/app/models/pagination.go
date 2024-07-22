package models

type PaginationResponse struct {
	CurrentPage int `json:"current_page"`
	PrevPage    int `json:"prev_page"`
	NextPage    int `json:"next_page"`
	MaxPage     int `json:"max_page"`
	Total       int `json:"total"`
}
