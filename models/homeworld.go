package models 

type Homeworld struct {
	ID             uint64 `json:"id"`
	URL            string `json:"url"`
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surface_water"`
	Population     string `json:"population"`
}
