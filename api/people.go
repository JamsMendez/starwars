package api

type People struct {
	ID        uint64 `json:"id"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hairColor"`
	SkinColor string `json:"sinkColor"`
	EyeColor  string `json:"eyeColor"`
	BirthYear string `json:"birthYear"`
	Gender    string `json:"gender"`
	Image     string `json:"image"`

	Homeworld Homeworld  `json:"Homeworld"`
	Films     []Film     `json:"films"`
	Vehicles  []Vehicle  `json:"vehicles"`
	Starships []Starship `json:"starships"`

	URLHomeworld string   `json:"-"`
	URLFilms     []string `json:"-"`
	URLVehicles  []string `json:"-"`
	URLStarships []string `json:"-"`
}

type Peoples struct {
	Results  []People `json:"results"`
	Next     uint64   `json:"next"`
	Previous uint64   `json:"previous"`
}
