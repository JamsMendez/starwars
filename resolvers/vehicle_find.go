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

// GetVehicleFind ...
func GetVehicleFind() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var vehicles []api.Vehicle

		var urls []string

		// People
		parentName := params.Info.ParentType.Name()
		if parentName != "" {
			if parentName == util.KeyPeople {
				peopleOne, isOk := params.Source.(api.People)
				if isOk && peopleOne.URL != "" {
					urls = peopleOne.URLVehicles
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

								vehicleJSON := models.Vehicle{}
								err = json.Unmarshal(buffer, &vehicleJSON)
								if err == nil {

									vehicle := api.Vehicle{
										ID:                   util.ParseURLToID(vehicleJSON.URL),
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
										VehicleClass:         vehicleJSON.VehicleClass,
									}

                  vehicle.Image = fmt.Sprintf("https://starwars-visualguide.com/assets/img/vehicles/%d.jpg", vehicle.ID)

									vehicles  = append(vehicles, vehicle)
								}
							}
						}
					}

					wg.Done()

				}(s, &wg)
			}
		}

		wg.Wait()

		sort.Slice(vehicles, func(i, j int) bool {
			return vehicles[i].ID < vehicles[j].ID
		})

		return vehicles, nil
	}
}
