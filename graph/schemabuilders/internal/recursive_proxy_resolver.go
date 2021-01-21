package internal

import (
	"bytes"
	"github.com/graphql-go/graphql"
)

type ProxyResolverContext struct {
	Resolver graphql.FieldResolveFn
	IsSealed bool
}

type GraphQLSubtreeReconstructorFn func(typeName string, arguments []string)

func NewProxyResolverContext() *ProxyResolverContext {
	return &ProxyResolverContext{
		Resolver: nil,
		IsSealed: true,
	}
}

func (prc *ProxyResolverContext) CreateResolver() *ProxyResolverContext {
	if prc.IsSealed {
		panic("cannot redeclare root resolver once sealed")
	}
	prc.IsSealed = true
	prc.Resolver = func(p graphql.ResolveParams) (interface{}, error) {
		// is root resolver
		var buf bytes.Buffer

		// query reconstructor
		return func(typeName string, arguments []string) {

		}, nil
	}
	return prc
}

func (prc *ProxyResolverContext) CreateSubtree(typeName string) *ProxyResolverContext {

	var resolver graphql.FieldResolveFn = func(p graphql.ResolveParams) (interface{}, error) {
		reconstruct, ok := p.Source.(GraphQLSubtreeReconstructorFn)
		if !ok {
			panic("source is not a reconstructor fn")
		}

		reconstruct(p.Info.FieldName, nil)
	}

	prc.

}
