package api

type Starship struct {
	ID                   uint64 `json:"id"`
	URL                  string `json:"url"`
	Name                 string `json:"name"`
	Model                string `json:"model"`
	Manufacture          string `json:"manufacture"`
	CostInCredits        string `json:"costInCredits"`
	Length               string `json:"length"`
	MaxAtmospheringSpeed string `json:"maxAtmospheringSpeed"`
	Crew                 string `json:"crew"`
	Passenger            string `json:"passenger"`
	CargoCapacity        string `json:"cargoCapacity"`
	Consumables          string `json:"consumables"`
	HyperdriveRating     string `json:"HyperdriveRating"`
	MGLT                 string `json:"mglt"`
	StarshipClass        string `json:"StarshipClass"`
	Image        string `json:"image"`
}
