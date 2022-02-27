package models

type ResponsePeople struct {
	Count    uint64   `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []People `json:"results"`
}
