package api

type Specie struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	AverageHeight   string `json:"averageHeight"`
	AverageLifespan string `json:"averageLifespan"`
	Designation     string `json:"designation"`
	EyeColors       string `json:"eyeColors"`
	HairColors      string `json:"hairColors"`
	SkinColors      string `json:"skinColors"`
	Language        string `json:"language"`
}
