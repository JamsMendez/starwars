package schemas

import (
	"github.com/graphql-go/graphql"

	"github.com/JamsMendez/starwars/util"
)

// FilmType ... Esquema del Film
var FilmType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: util.KeyFilm,
		Fields: graphql.Fields{
			util.KeyFieldID:          &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			util.KeyFieldURL:         &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldEpisodeID:   &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			util.KeyFieldTitle:       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldOpenCrawl:   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldDirector:    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldProducer:    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			util.KeyFieldReleaseDate: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
