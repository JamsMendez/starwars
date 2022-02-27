package resolvers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JamsMendez/starwars/api"
	"github.com/JamsMendez/starwars/models"
	"github.com/JamsMendez/starwars/util"
	"github.com/graphql-go/graphql"
)

// GetStarshipFind ...
func GetStarshipFind() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var vehicles []api.Starship

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

		for _, s := range urls {
			rUrl, err := url.Parse(s)
			if err == nil {

				response, err := http.Get(rUrl.String())
				if err == nil {

					defer response.Body.Close()
					buffer, err := ioutil.ReadAll(response.Body)
					if err == nil {

						vehicleJSON := models.Starship{}
						err = json.Unmarshal(buffer, &vehicleJSON)
						if err == nil {

							vehicle := api.Starship{
								URL:                  vehicleJSON.URL,
								Name:                 vehicleJSON.Name,
								Model:                vehicleJSON.Model,
								Manufacture:          vehicleJSON.Manufacture,
								CostInCredits:        vehicleJSON.CostInCredits,
								Length:               vehicleJSON.Length,
								MaxAtmospheringSpeed: vehicleJSON.URL,
								Crew:                 vehicleJSON.Crew,
								Passenger:            vehicleJSON.Passenger,
								CargoCapacity:        vehicleJSON.CargoCapacity,
								Consumables:          vehicleJSON.Consumables,
								HyperdriveRating:     vehicleJSON.HyperdriveRating,
								MGLT:                 vehicleJSON.MGLT,
								StarshipClass:        vehicleJSON.StarshipClass,
							}

							vehicles = append(vehicles, vehicle)
						}
					}
				}
			}
		}

		return vehicles, nil
	}
}
