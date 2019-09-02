package graphqlscalar

import (
	"mime/multipart"

	"github.com/graphql-go/graphql"
)

// Files upload struct
type Files struct {
	value []*multipart.FileHeader
}

func (f *Files) fileValue() []*multipart.FileHeader {
	return f.value
}

// MultipartScalarType custom scalar type Multipart
var MultipartScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "MultipartScalarType",
	Description: "The `MultipartScalarType` scalar type represents an multipart file",
	// Serialize serializes `CustomID` to multipart.FileHeader
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case Files:
			return value.fileValue()
		case *Files:
			v := *value
			return v.fileValue()
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `CustomID`.
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case Files:
			return value.fileValue()
		case *Files:
			v := *value
			return v.fileValue()
		default:
			return nil
		}
	},
})
