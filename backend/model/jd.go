package model

type JDRequest struct {
	Position         string   `json:"position"`
	Location         string   `json:"location"`
	Responsibilities []string `json:"responsibilities"`
	Requirements     []string `json:"requirements"`
}
