package api

type Vehicle struct {
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
	VehicleClass         string `json:"vehicleClass"`
}
