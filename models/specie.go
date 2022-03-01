package models

type Specie struct {
	ID              uint64 `json:"id"`
	URL             string `json:"url"`
	Name            string `json:"name"`
	AverageHeight   string `json:"average_height"`
	AverageLifespan string `json:"average_lifespan"`
	Classification  string `json:"classification"`
	Designation     string `json:"designation"`
	EyeColors       string `json:"eye_colors"`
	HairColors      string `json:"hair_colors"`
	Language        string `json:"language"`
	SkinColors      string `json:"skin_colors"`
}
