package schemabuilders

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle-sdk/graph"
	. "github.com/terra-project/mantle-sdk/graph/schemabuilders/internal"
)

type (
	RemoteQueriesMap map[string]TypeDescriptor
	ResolverCreator  func(name string) graphql.FieldResolveFn
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

		reconstructQuery(
			rootQuery,
			fields,
			remoteModelsMap,
			func(name string) graphql.FieldResolveFn {
				return func(p graphql.ResolveParams) (interface{}, error) {

					// if root, pass down a function which has

					// else
					return func() {

						sgrf, isRoot := p.Source.(SubGraphRecFunc)

					}, nil
				}
			},
		)

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

// reconstructQuery builds remote schemas available in RootQuery
func reconstructQuery(
	td TypeDescriptor,
	targetFields *graphql.Fields,
	rqm RemoteQueriesMap,
	resolverCreator func(name string) graphql.FieldResolveFn,
) {
	var fields = graphql.Fields{}
	for _, f := range td.Fields {
		fields[f.Name] = &graphql.Field{
			Name:              f.Name,
			Type:              reconstructField(f),
			Args:              reconstructFieldArgs(f),
			Resolve:           nil,
			DeprecationReason: "",
			Description:       "",
		}
	}

	query := &graphql.Field{
		Name:              td.Name,
		Type:              nil,
		Args:              nil,
		Resolve:           nil,
		DeprecationReason: "",
		Description:       td.Description,
	}

}

func reconstructField(
	f Field,
) graphql.Output {

}

func reconstructFieldArgs(
	f Field,
) graphql.FieldConfigArgument {
	argumentConfig := graphql.FieldConfigArgument{}
	for _, arg := range f.Args {

		argumentConfig[arg.Name] = &graphql.ArgumentConfig{
			Type:         ,
			DefaultValue: arg.DefaultValue,
			Description:  arg.Description,
		}
	}
}

func reconstructType(
	t Type,
) {
	t.
}