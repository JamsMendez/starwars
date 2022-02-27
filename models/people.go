package models

type People struct {
	ID        uint64 `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"sink_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`

	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
}
