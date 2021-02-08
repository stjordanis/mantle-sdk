package schemabuilders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle-sdk/graph"
	. "github.com/terra-project/mantle-sdk/graph/schemabuilders/proxy_resolver"
	"io/ioutil"
	"net/http"
)

type (
	RemoteQueriesMap map[string]TypeDescriptor
	SubGraphRecFunc  func(nodeName string)
)

var (
	errInvalidContext = fmt.Errorf("context is not proxy resolver context")
	errInvalidSource  = fmt.Errorf("source is not proxy resolver context")
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

		// create rootProxyResolverContext
		// once query building is done, the callback will relay the graphql query
		// back tothe base mantle (to remoteMantleEndpoint)
		rootProxyResolverContext := NewProxyResolverContext(func(query []byte) map[string]interface{} {
			response := graph.CreateRemoteMantleRequest(remoteMantleEndpoint, query)

			gqlResponse := new(struct {
				Data map[string]interface{} `json:"data"`
			})

			err := json.Unmarshal(response, gqlResponse)
			if err != nil {
				panic(err)
			}

			return gqlResponse.Data
		})

		// iterate over queriable field, reconstruct query
		for _, queriableField := range rootQueryFields {
			name := queriableField.Name

			// query output object
			fieldOutput := reconstructFieldConfig(
				queriableField.Type,
				remoteModelsMap,
			)

			fieldArguments := reconstructFieldArgument(
				queriableField.Args,
				remoteModelsMap,
			)

			//
			(*fields)[name] = &graphql.Field{
				Name: name,
				Type: fieldOutput,
				Args: fieldArguments,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return rootProxyResolverContext.CreateSubtree(name, p.Args), nil
				},
				DeprecationReason: fmt.Sprintf("(proxy) %s", queriableField.DeprecationReason),
				Description:       fmt.Sprintf("(proxy) %s", queriableField.Description),
			}
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
			_, isScalarType := selectionType.(*graphql.Scalar)

			subselectionFields[selectionName] = &graphql.Field{
				Name: selectionName,
				Type: selectionType,
				Args: nil,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// have a reference to ProxyResolver, depending on the tree order
					var prc *ProxyResolverContext
					prc, ok := p.Source.(*ProxyResolverContext)
					if !ok {
						return nil, errInvalidSource
					}

					var isLeafNode = isScalarType
					if isLeafNode {
						return prc.Resolve()
					} else {
						return prc.CreateSubtree(selectionName, p.Args), nil
					}
				},
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

func reconstructFieldArgument(
	queryArguments []Argument,
	remoteQueriesMap RemoteQueriesMap,
) graphql.FieldConfigArgument {
	argumentConfig := graphql.FieldConfigArgument{}

	for _, queryArgument := range queryArguments {

	}

	return argumentConfig
}
