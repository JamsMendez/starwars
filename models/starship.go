package models

type Starship struct {
	ID                   uint64 `json:"id"`
	URL                  string `json:"url"`
	Name                 string `json:"name"`
	Model                string `json:"model"`
	Manufacture          string `json:"manufacture"`
	CostInCredits        string `json:"cost_in_credits"`
	Length               string `json:"length"`
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
	Crew                 string `json:"crew"`
	Passenger            string `json:"passenger"`
	CargoCapacity        string `json:"cargo_capacity"`
	Consumables          string `json:"consumables"`
	HyperdriveRating     string `json:"hyperdrive_rating"`
	MGLT                 string `json:"MGLT"`
	StarshipClass        string `json:"starship_class"`
}
