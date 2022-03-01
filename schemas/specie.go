package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/util"
)

// SpecieType ... Esquema del Specie
var SpecieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: util.KeySpecie,
		Fields: graphql.Fields{
			util.KeyFieldID:              &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			util.KeyFieldURL:             &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldName:            &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldAverageHeight:   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldAverageLifespan: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldHairColors:      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldSkinColors:      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldEyeColors:       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldClassification:  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldDesignation:     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldLanguage:        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
