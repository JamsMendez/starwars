package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/resolvers"
	"github.com/JamsMendez/starwars/util"
)

// QueryHomeworld...
func QueryHomeworld() *graphql.Field {
	return &graphql.Field{
		Type:    HomeworldType,
		Resolve: resolvers.GetHomeworldFindOne(),
	}
}

// QueryPeople ...
func QueryPeople() *graphql.Field {
	return &graphql.Field{
		Type:    PeopleType,
		Args: graphql.FieldConfigArgument{
			util.KeyFieldID: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: resolvers.GetPeopleFindOne(),
	}
}

// QueryPeoples ...
func QueryPeoples() *graphql.Field {
	return &graphql.Field{
		Type:    PeoplesType,
		Args: graphql.FieldConfigArgument{
			util.KeyFieldPage: &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			util.KeyFieldSearch: &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: resolvers.GetPeopleFind(),
	}
}

// QueryFilms...
func QueryFilms() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(FilmType),
		Resolve: resolvers.GetFilmFind(),
	}
}

// QueryVehicles...
func QueryVehicles() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(VehicleType),
		Resolve: resolvers.GetVehicleFind(),
	}
}

// QueryStarships...
func QueryStarships() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(StarshipType),
		Resolve: resolvers.GetStarshipFind(),
	}
}
