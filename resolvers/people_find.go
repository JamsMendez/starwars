package resolvers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/api"
	"github.com/JamsMendez/starwars/models"
	"github.com/JamsMendez/starwars/util"
)

// GetPeopleFind ...
func GetPeopleFind() graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		var peoples api.Peoples

		var page uint64

		if value, isOk := params.Args[util.KeyFieldPage]; isOk {
			if valueInt, isOk := value.(int); isOk {
				if isOk {
					page = uint64(valueInt)
				}
			}
		}

		s := fmt.Sprintf("%s%s", util.ApiBase, util.ApiPeople)
		rUrl, err := url.Parse(s)
		if err != nil {
			return peoples, nil
		}

		if page > 0 {
      q := rUrl.Query()
      q.Set(util.KeyFieldPage, fmt.Sprintf("%d", page))
      rUrl.RawQuery = q.Encode()
		}

		response, err := http.Get(rUrl.String())
		if err != nil {
			return peoples, nil
		}

		defer response.Body.Close()
		buffer, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return peoples, nil
		}

		responseJSON := models.ResponsePeople{}
		err = json.Unmarshal(buffer, &responseJSON)
		if err != nil {
			return peoples, nil
		}

		peoplesJSON := responseJSON.Results

		var results []api.People

		for _, peopleJSON := range peoplesJSON {
			people := api.People{
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
				URLHomeworld: peopleJSON.Homeworld,
			}

			results = append(results, people)
		}

		peoples.Results = results

		if responseJSON.Next != "" {
			nextUrl, err := url.Parse(responseJSON.Next)
			if err == nil {
				queries := nextUrl.Query()
				value := queries.Get(util.KeyFieldPage)
				if value != "" {
					nextInt, err := strconv.Atoi(value)
					if err == nil {
						peoples.Next = uint64(nextInt)
					}
				}
			}
		}

		if responseJSON.Previous != "" {
			nextUrl, err := url.Parse(responseJSON.Previous)
			if err == nil {
				queries := nextUrl.Query()
				value := queries.Get(util.KeyFieldPage)
				if value != "" {
					previousInt, err := strconv.Atoi(value)
					if err == nil {
						peoples.Previous = uint64(previousInt)
					}
				}
			}
		}

		return peoples, nil
	}
}
