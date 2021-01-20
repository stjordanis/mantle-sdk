package internal

import (
	"github.com/graphql-go/graphql"
)

type ProxyResolverContext struct {
	Resolver graphql.FieldResolveFn
	IsSealed bool
	Children []graphql.FieldResolveFn
}

func NewProxyResolverContext() *ProxyResolverContext {
	return &ProxyResolverContext{
		Resolver: nil,
		Children: make([]graphql.FieldResolveFn, 0),
	}
}

func (prc *ProxyResolverContext) CreateRootQuery() *ProxyResolverContext {
	if prc.IsSealed {
		panic("cannot redeclare root resolver once sealed")
	}
	prc.IsSealed = true
	prc.Resolver = func(p graphql.ResolveParams) (interface{}, error) {

	}
	return prc
}

func (prc *ProxyResolverContext) CreateSubtree(typeName string) {

}
