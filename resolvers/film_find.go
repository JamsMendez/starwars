package resolvers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/JamsMendez/starwars/api"
	"github.com/JamsMendez/starwars/models"
	"github.com/JamsMendez/starwars/util"
	"github.com/graphql-go/graphql"
)

// GetFilmFind ...
func GetFilmFind() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var films []api.Film

		var urls []string

		// People
		parentName := params.Info.ParentType.Name()
		if parentName != "" {
			if parentName == util.KeyPeople {
				peopleOne, isOk := params.Source.(api.People)
				if isOk && peopleOne.URL != "" {
					urls = peopleOne.URLFilms
				}
			}
		}

		var wg sync.WaitGroup
		size := len(urls)
		if size > 0 {
			wg.Add(size)

			for _, s := range urls {
				go func(s string, wg *sync.WaitGroup) {
					rUrl, err := url.Parse(s)
					if err == nil {

						response, err := http.Get(rUrl.String())
						if err == nil {

							defer response.Body.Close()
							buffer, err := ioutil.ReadAll(response.Body)
							if err == nil {

								filmJSON := models.Film{}
								err = json.Unmarshal(buffer, &filmJSON)
								if err == nil {

									film := api.Film{
										URL:          filmJSON.URL,
										EpisodeID:    filmJSON.EpisodeID,
										Title:        filmJSON.Title,
										OpeningCrawl: filmJSON.OpeningCrawl,
										Director:     filmJSON.Director,
										Producer:     filmJSON.Producer,
										ReleaseDate:  filmJSON.ReleaseDate,
									}

									films = append(films, film)
								}
							}
						}
					}

					wg.Done()

				}(s, &wg)
			}
		}

		wg.Wait()

		return films, nil
	}
}
