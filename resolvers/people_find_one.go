package resolvers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/api"
	"github.com/JamsMendez/starwars/models"
	"github.com/JamsMendez/starwars/util"
)

// GetPeopleFindOne ...
func GetPeopleFindOne() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var people api.People

		var ID uint64

		if value, isOk := params.Args[util.KeyFieldID]; isOk {
			if valueInt, isOk := value.(int); isOk {
				if isOk {
					ID = uint64(valueInt)
				}
			}
		}

		s := fmt.Sprintf("%s%s/%d/", util.ApiBase, util.ApiPeople, ID)
		rUrl, err := url.Parse(s)
		if err != nil {
			return nil, nil
		}

		response, err := http.Get(rUrl.String())
		if err != nil {
			return nil, nil
		}

		defer response.Body.Close()
		buffer, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, nil
		}

		peopleJSON := models.People{}
		err = json.Unmarshal(buffer, &peopleJSON)
		if err != nil {
			return nil, nil
		}

		people = api.People{
			ID:        util.ParseURLToID(peopleJSON.URL),
			URL:       peopleJSON.URL,
			Name:      peopleJSON.Name,
			Height:    peopleJSON.Height,
			Mass:      peopleJSON.Mass,
			HairColor: peopleJSON.HairColor,
			SkinColor: peopleJSON.SkinColor,
			EyeColor:  peopleJSON.EyeColor,
			BirthYear: peopleJSON.BirthYear,
			Gender:    peopleJSON.Gender,

			URLFilms:     peopleJSON.Films,
			URLVehicles:  peopleJSON.Vehicles,
			URLStarships: peopleJSON.Starships,
			URLSpecies:   peopleJSON.Species,
			URLHomeworld: peopleJSON.Homeworld,
		}

		people.Image = fmt.Sprintf("https://starwars-visualguide.com/assets/img/characters/%d.jpg", people.ID)

		return people, nil
	}
}
