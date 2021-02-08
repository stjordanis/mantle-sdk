package proxy_resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle-sdk/graph/generate"
)

type (
	Definitions map[string]TypeDescriptor
)

func GetArgumentGraphQLType(argType *Input, definitions Definitions) graphql.Input {
	underlyingType, ok := definitions[argType.Name]

	if !ok {
		panic(fmt.Errorf("underlying type %s does not exist in definitions map", argType.Name))
	}

	switch underlyingType.Kind {
	// graphql-native scalars + custom scalars
	case "SCALAR":
		switch underlyingType.Name {
		case "Int":
			return graphql.Int
		case "Float":
			return graphql.Float
		case "String":
			return graphql.String
		case "Boolean":
			return graphql.Boolean
		case "ID":
			return graphql.ID
		case "DateTime":
			return graphql.DateTime
		default:
			// check cosmos-scalar map
			cosmosScalar := generate.GetCosmosScalarByName(underlyingType.Name)
			if cosmosScalar == nil {
				panic(fmt.Errorf("custom scalar %s not found", underlyingType.Name))
			}

			return cosmosScalar
		}

	// get the matching OBJECT type of definitions, reconstruct them into graphql type
	case "OBJECT":

	case "LIST":

	default:

	}
}
