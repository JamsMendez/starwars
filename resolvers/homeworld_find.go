package resolvers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JamsMendez/starwars/api"
	"github.com/JamsMendez/starwars/models"
	"github.com/JamsMendez/starwars/util"
	"github.com/graphql-go/graphql"
)

// GetHomeworldFindone ...
func GetHomeworldFindOne() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var homeworld api.Homeworld

		var s string

		// People
		parentName := params.Info.ParentType.Name()
		if parentName != "" {
			if parentName == util.KeyPeople {
				peopleOne, isOk := params.Source.(api.People)
				if isOk && peopleOne.URL != "" {
					s = peopleOne.URLHomeworld
				}
			}
		}

		rUrl, err := url.Parse(s)
		if err == nil {

			response, err := http.Get(rUrl.String())
			if err == nil {

				defer response.Body.Close()
				buffer, err := ioutil.ReadAll(response.Body)
				if err == nil {

					homeworldJSON := models.Homeworld{}
					err = json.Unmarshal(buffer, &homeworldJSON)
					if err == nil {

						homeworld = api.Homeworld{
							ID:             util.ParseURLToID(homeworldJSON.URL),
							URL:            homeworldJSON.URL,
							Name:           homeworldJSON.Name,
							RotationPeriod: homeworldJSON.RotationPeriod,
							OrbitalPeriod:  homeworldJSON.OrbitalPeriod,
							Diameter:       homeworldJSON.Diameter,
							Climate:        homeworldJSON.Climate,
							Gravity:        homeworldJSON.Gravity,
							Terrain:        homeworldJSON.Terrain,
							SurfaceWater:   homeworldJSON.SurfaceWater,
							Population:     homeworldJSON.Population,
						}

            homeworld.Image = fmt.Sprintf("https://starwars-visualguide.com/assets/img/planets/%d.jpg", homeworld.ID)
					}
				}
			}
		}

		return homeworld, nil
	}
}
