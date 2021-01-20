package schemabuilders

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle-sdk/graph"
	. "github.com/terra-project/mantle-sdk/graph/schemabuilders/internal"
)

type (
	RemoteQueriesMap map[string]TypeDescriptor
	SubGraphRecFunc  func(nodeName string)
)

func CreateRemoteModelSchemaBuilder(remoteMantleEndpoint string) graph.SchemaBuilder {
	return func(fields *graphql.Fields) error {
		schema := NewIntrospection(remoteMantleEndpoint)

		// 1. go through all types, and create objects first
		remoteModelsMap := convertTypesToTypeMap(schema.Types)

		// 2. iterate through all FIELDS in RootQuery,
		//    and map out all things recursively
		rootQuery, ok := remoteModelsMap["RootQuery"]
		if !ok {
			return fmt.Errorf("remote mantle does not have root query")
		}

		rootQueryFields := rootQuery.Fields

		// iterate over queriable field, reconstruct query
		for _, queriableField := range rootQueryFields {
			name := queriableField.Name
			queryType := queriableField.Type

			// query output object
			fieldConfig := reconstructFieldConfig(
				queryType,
				remoteModelsMap,
			)

			(*fields)[name] = fieldConfig
		}

		return nil
	}
}

func convertTypesToTypeMap(types []TypeDescriptor) RemoteQueriesMap {
	m := make(map[string]TypeDescriptor)

	for _, t := range types {
		m[t.Name] = t
	}

	return m
}

func reconstructFieldConfig(
	queryType Type,
	remoteQueriesMap RemoteQueriesMap,
) graphql.Output {
	switch queryType.Kind {
	case "OBJECT":
		outputTypeName, ok := queryType.Name.(string)
		if !ok {
			panic(fmt.Errorf("output type name is not string; trying to reconstructFieldConfig %v", queryType))
		}
		definition, ok := remoteQueriesMap[outputTypeName]
		if !ok {
			panic(fmt.Errorf("schema definition %s does not exist", outputTypeName))
		}

		subselectionFields := graphql.Fields{}

		subSelections := definition.Fields
		for _, selection := range subSelections {
			selectionName := selection.Name
			selectionType := reconstructFieldConfig(
				selection.Type,
				remoteQueriesMap,
			)

			subselectionFields[selectionName] = &graphql.Field{
				Name:              selectionName,
				Type:              selectionType,
				Args:              nil,
				Resolve:           nil,
				DeprecationReason: selection.DeprecationReason,
				Description:       selection.Description,
			}
		}

		return graphql.NewObject(graphql.ObjectConfig{
			Name:        outputTypeName,
			Fields:      subselectionFields,
			Description: definition.Description,
		})

	case "LIST":

	default:

	}
}
