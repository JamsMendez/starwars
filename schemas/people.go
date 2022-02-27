package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/util"
)

// PeopleType ... Esquema del People
var PeopleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: util.KeyPeople,
		Fields: graphql.Fields{
			util.KeyFieldID:        &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			util.KeyFieldURL:       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldName:      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldHeight:    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldMass:      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldSkinColor: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldEyeColor:  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldBirthYear: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldGender:    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},

			util.KeyFieldHomeworld: QueryHomeworld(),
			util.KeyFieldFilms:     QueryFilms(),
			util.KeyFieldVehicles:  QueryVehicles(),
			util.KeyFieldStarships: QueryStarships(),
		},
	},
)

// PeoplesType ... Esquema del Peoples
var PeoplesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: util.KeyPeoples,
		Fields: graphql.Fields{
			util.KeyFieldResults:  &graphql.Field{Type: graphql.NewList(PeopleType)},
			util.KeyFieldNext:     &graphql.Field{Type: graphql.Int},
			util.KeyFieldPrevious: &graphql.Field{Type: graphql.Int},
		},
	},
)
