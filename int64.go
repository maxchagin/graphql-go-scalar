package graphqlscalar

import (
	"math"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func coerceInt64(value interface{}) interface{} {
	switch value := value.(type) {
	case int64:
		return value
	case *int64:
		return *value
	case int:
		return int64(value)
	case *int:
		return int64(*value)
	case float64:
		if value < float64(math.MinInt64) || value > float64(math.MaxInt64) {
			return nil
		}
		return int64(value)
	case *float64:
		if value == nil {
			return nil
		}
		return int64(*value)
	case string:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil
		}
		return val
	case *string:
		if value == nil {
			return nil
		}
		val, err := strconv.ParseInt(*value, 10, 64)
		if err != nil {
			return nil
		}
		return val
	default:
		return nil
	}
}

// Int64 кастомный тип для int64
var Int64 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Int64",
	Description: "The `Int64` scalar type for int64",
	// Serialize serializes `CustomID` to multipart.FileHeader
	Serialize: coerceInt64,
	// ParseValue parses GraphQL variables from `int64` to `Int64`.
	ParseValue: coerceInt64,
	// ParseLiteral parses GraphQL AST value to `CustomInt64`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if int64Value, err := strconv.ParseInt(valueAST.Value, 10, 64); err == nil {
				return int64Value
			}
		}
		return nil
	},
})
