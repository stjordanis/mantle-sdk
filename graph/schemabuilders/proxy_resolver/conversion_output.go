package proxy_resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle-sdk/graph/generate"
)

func GetGraphQLOutputType(argType *Input, definitions Definitions) graphql.Input {
	underlyingType, ok := definitions[argType.Name]

	if !ok {
		panic(fmt.Errorf("underlying type %s does not exist in definitions map", argType.Name))
	}

	switch underlyingType.Kind {
	// graphql-native scalars + custom scalars
	case "SCALAR":
		switch underlyingType.Name {
		// take care of all graphql-native scalars
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
		// check cosmos-scalar map
		default:
			cosmosScalar := generate.GetCosmosScalarByName(underlyingType.Name)

			// if name is unknown, mantle can't handle it. panic here
			if cosmosScalar == nil {
				panic(errUnknownCustomScalar(underlyingType.Name))
			}

			return cosmosScalar
		}

	// get the matching OBJECT type of definitions, reconstruct them into graphql type
	case "OBJECT":
		panic(fmt.Errorf("object argument is not supported yet"))

	case "LIST":
		ofType := argType.OfType
		underlyingType := GetGraphQLOutputType(ofType, definitions)
		return graphql.NewList(underlyingType)

	// unknown type, panic
	default:
		panic(errUnknownArgumentType(argType.Name))
	}
}
