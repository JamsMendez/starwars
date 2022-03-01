package resolvers

import (
	"encoding/json"
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

// GetSpecieFind ...
func GetSpecieFind() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var species []api.Specie

		var urls []string

		// People
		parentName := params.Info.ParentType.Name()
		if parentName != "" {
			if parentName == util.KeyPeople {
				peopleOne, isOk := params.Source.(api.People)
				if isOk && peopleOne.URL != "" {
					urls = peopleOne.URLSpecies
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

								specieJSON := models.Specie{}
								err = json.Unmarshal(buffer, &specieJSON)
								if err == nil {

									specie := api.Specie{
										ID:              util.ParseURLToID(specieJSON.URL),
										Name:            specieJSON.Name,
										URL:             specieJSON.URL,
										AverageHeight:   specieJSON.AverageHeight,
										AverageLifespan: specieJSON.AverageLifespan,
										Designation:     specieJSON.Designation,
										EyeColors:       specieJSON.EyeColors,
										HairColors:      specieJSON.HairColors,
										Language:        specieJSON.Language,
									}

									species = append(species, specie)
								}
							}
						}
					}

					wg.Done()

				}(s, &wg)
			}
		}

		wg.Wait()

		sort.Slice(species, func(i, j int) bool {
			return species[i].ID < species[j].ID
		})

		return species, nil
	}
}
