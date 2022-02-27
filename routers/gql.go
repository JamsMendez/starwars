package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"

	"github.com/JamsMendez/starwars/schemas"
	"github.com/JamsMendez/starwars/util"
)

// GQL ...
type GQL struct {
	API *echo.Group
}

// New ...
func (g GQL) New() {
	// Schema
	schema, err := graphql.NewSchema(schemas.SchemaConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	g.API.GET("/graphql", func(c echo.Context) error {
		query := c.QueryParam(util.KeyQuery)
		variablesIn := c.QueryParam(util.KeyVariables)
		operation := c.QueryParam(util.KeyOperation)

		variables := map[string]interface{}{}
		if variablesIn != "" {
			buffer := []byte(variablesIn)
			err := json.Unmarshal(buffer, &variables)
			if err != nil {
				fmt.Println("GQL.Variables.Unmarshal.: ", err)

				return c.NoContent(http.StatusBadRequest)
			}
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  query,
			VariableValues: variables,
			OperationName:  operation,
			RootObject:     map[string]interface{}{},
		})

		if result.HasErrors() {
			return c.JSON(http.StatusAccepted, result.Errors)
		}

		return c.JSON(http.StatusOK, result.Data)

	})
}
