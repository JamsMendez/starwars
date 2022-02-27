package routers

import (
	"fmt"

	"github.com/JamsMendez/starwars/schemas"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
)

// GQLDev ...
type GQLDev struct {
	API *echo.Group
}

// New ...
func (g GQLDev) New() {
	// Schema
	schema, err := graphql.NewSchema(schemas.SchemaConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	g.API.GET("/graphql/dev", echo.WrapHandler(h))
	g.API.POST("/graphql/dev", echo.WrapHandler(h))
}
