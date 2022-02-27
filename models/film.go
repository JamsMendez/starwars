package models 

type Film struct {
	ID           uint64 `json:"id"`
	EpisodeID    uint64 `json:"episode_id"`
	URL          string `json:"url"`
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
	ReleaseDate  string `json:"release_date"`
}
