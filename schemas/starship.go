package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/util"
)

// StarshipType ... Esquema del Starship
var StarshipType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: util.KeyStarship,
		Fields: graphql.Fields{
			util.KeyFieldID:                   &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			util.KeyFieldURL:                  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldName:                 &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldModel:                &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldManufacture:          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldCostInCredit:         &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldLength:               &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldMaxAtmospheringSpeed: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldCrew:                 &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldPassenger:            &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldCargoCapacity:        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldConsumables:          &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldHyperdriveRating:     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldMGLT:                 &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldStarshipClass:        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
