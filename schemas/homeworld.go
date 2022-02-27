package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/util"
)

// HomeworldType ... Esquema del Homeworld
var HomeworldType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: util.KeyHomeworld,
		Fields: graphql.Fields{
			util.KeyFieldID:             &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			util.KeyFieldURL:            &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldName:           &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldRotationPeriod: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldOrbitalPeriod:  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldDiameter:       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldClimate:        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldGravity:        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldTerrain:        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldSurfaceWater:   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldPopulation:     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldImage:          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
