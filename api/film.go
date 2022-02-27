package api

type Film struct {
	ID           uint64 `json:"id"`
	EpisodeID    uint64 `json:"episodeId"`
	URL          string `json:"url"`
	Title        string `json:"title"`
	OpeningCrawl string `json:"openingCrawl"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
	ReleaseDate  string `json:"releaseDate"`
}
