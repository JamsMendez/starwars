package util

const (
	KeyQuery     = "query"
	KeyVariables = "variables"
	KeyOperation = "operation"

	KeyJSON   = "json"
	KeyBuffer = "buffer"

	NA = "N/A"
)

// GQLIn ...
type GQLIn struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
	Operation string                 `json:"operation"`
}

// MsgsError ...
type MsgsError struct {
	Messages []string `json:"msgs"`
}

// ResJSON JSON
type ResJSON struct {
	O interface{} `json:"data"`
}

// ResJSONs []JSON
type ResJSONs struct {
	Os interface{} `json:"data"`
}
