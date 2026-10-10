package graphqlscalar

import (
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// Timestamp custom scalar type int64
var Timestamp = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Timestamp",
	Description: "The `Timestamp` custom scalar type for time.Time. Represents a Unix timestamp in seconds since the epoch. All values are strictly parsed, normalized, and returned in the UTC timezone.",
	Serialize: func(value interface{}) interface{} {
		switch v := value.(type) {
		case time.Time:
			return v.Unix()
		case *time.Time:
			if v == nil {
				return nil
			}
			return v.Unix()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch v := value.(type) {
		case int64:
			return time.Unix(v, 0).UTC()
		case int:
			return time.Unix(int64(v), 0).UTC()
		case float64:
			return time.Unix(int64(v), 0).UTC()
		case string:
			if timestamp, err := strconv.ParseInt(v, 10, 64); err == nil {
				return time.Unix(timestamp, 0).UTC()
			}
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				return t.UTC()
			}
			return nil
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch v := valueAST.(type) {
		case *ast.IntValue:
			if val, err := strconv.ParseInt(v.Value, 10, 64); err == nil {
				return time.Unix(val, 0).UTC()
			}
			return nil
		case *ast.StringValue:
			if timestamp, err := strconv.ParseInt(v.Value, 10, 64); err == nil {
				return time.Unix(timestamp, 0).UTC()
			}
			if t, err := time.Parse(time.RFC3339, v.Value); err == nil {
				return t.UTC()
			}
			return nil
		case *ast.FloatValue:
			if val, err := strconv.ParseFloat(v.Value, 64); err == nil {
				return time.Unix(int64(val), 0).UTC()
			}
			return nil
		default:
			return nil
		}
	},
})
