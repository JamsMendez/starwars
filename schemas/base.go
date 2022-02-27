package schemas

import "github.com/graphql-go/graphql"

const (
	KeyQueryPeoples = "peoples"
	KeyQueryPeople = "people"
)

// Queries
var rootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		KeyQueryPeoples: QueryPeoples(),
		KeyQueryPeople: QueryPeople(),
	},
}

// Mutations
var rootMutation = graphql.ObjectConfig{
	Name:   "RootMutation",
	Fields: graphql.Fields{},
}

// SchemaConfig ...
var SchemaConfig = graphql.SchemaConfig{
	Query:    graphql.NewObject(rootQuery),
	//Mutation: graphql.NewObject(rootMutation),
}
