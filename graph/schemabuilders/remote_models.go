package schemabuilders

import (
	"bytes"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle-sdk/graph"
	. "github.com/terra-project/mantle-sdk/graph/schemabuilders/internal"
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
		rootProxyResolverContext := NewProxyResolverContext()

		// iterate over queriable field, reconstruct query
		for _, queriableField := range rootQueryFields {
			name := queriableField.Name
			queryType := queriableField.Type

			// query output object
			fieldConfig := reconstructFieldConfig(
				queryType,
				remoteModelsMap,
				rootProxyResolverContext,
				true,
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
	isRoot bool,
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
				false,
			)

			subselectionFields[selectionName] = &graphql.Field{
				Name: selectionName,
				Type: selectionType,
				Args: nil,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// have a reference to ProxyResolver, depending on the tree order
					var prc *ProxyResolverContext
					if isRoot {
						prc, ok = p.Context.Value(ProxyResolverContextKey).(*ProxyResolverContext)
						if !ok {
							panic(errInvalidContext)
						}
					} else {
						prc, ok = p.Source.(*ProxyResolverContext)
						if !ok {
							panic(errInvalidSource)
						}
					}

					//

					prc, ok := p.Source.(*ProxyResolverResponseCallback)
					if !ok {
						panic(fmt.Errorf("subselection %s source is not a proxy resolver context", selectionName))
					}

					cprc.AssignArguments(p.Args)

					if isRoot {
						rootQuery := ReconstructRootQuery(bytes.Buffer{}, cprc)

						// somehow make query
						response, err := http.Get("https://tequila-mantle.terra.dev", rootQuery.String())
						if err != nil {
							return nil, err
						}

						responseBuf, err := ioutil.ReadAll(response.Body)
						cprc.SetResponse(string(responseBuf), err)

						cprc.SetResponse()
					}

					return func() (interface{}, error) {
						return cprc.Resolve()
					}, nil
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
