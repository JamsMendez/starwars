package resolvers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"sync"

	"github.com/JamsMendez/starwars/api"
	"github.com/JamsMendez/starwars/models"
	"github.com/JamsMendez/starwars/util"
	"github.com/graphql-go/graphql"
)

// GetStarshipFind ...
func GetStarshipFind() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var starships []api.Starship

		var urls []string

		// People
		parentName := params.Info.ParentType.Name()
		if parentName != "" {
			if parentName == util.KeyPeople {
				peopleOne, isOk := params.Source.(api.People)
				if isOk && peopleOne.URL != "" {
					urls = peopleOne.URLStarships
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

								starshipJSON := models.Starship{}
								err = json.Unmarshal(buffer, &starshipJSON)
								if err == nil {

									starship := api.Starship{
										ID:                   util.ParseURLToID(starshipJSON.URL),
										URL:                  starshipJSON.URL,
										Name:                 starshipJSON.Name,
										Model:                starshipJSON.Model,
										Manufacture:          starshipJSON.Manufacture,
										CostInCredits:        starshipJSON.CostInCredits,
										Length:               starshipJSON.Length,
										MaxAtmospheringSpeed: starshipJSON.URL,
										Crew:                 starshipJSON.Crew,
										Passenger:            starshipJSON.Passenger,
										CargoCapacity:        starshipJSON.CargoCapacity,
										Consumables:          starshipJSON.Consumables,
										HyperdriveRating:     starshipJSON.HyperdriveRating,
										MGLT:                 starshipJSON.MGLT,
										StarshipClass:        starshipJSON.StarshipClass,
									}

                  starship.Image = fmt.Sprintf("https://starwars-visualguide.com/assets/img/starships/%d.jpg", starship.ID)

									starships = append(starships, starship)
								}
							}
						}
					}

					wg.Done()

				}(s, &wg)
			}
		}

		wg.Wait()

		sort.Slice(starships, func(i, j int) bool {
			return starships[i].ID < starships[j].ID
		})

		return starships, nil
	}
}
