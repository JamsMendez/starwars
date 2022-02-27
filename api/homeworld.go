package api

type Homeworld struct {
	ID             uint64 `json:"id"`
	URL            string `json:"url"`
	Name           string `json:"name"`
	RotationPeriod string `json:"rotationPeriod"`
	OrbitalPeriod  string `json:"orbitalPeriod"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surfaceWater"`
	Population     string `json:"population"`
	Image          string `json:"image"`
}
