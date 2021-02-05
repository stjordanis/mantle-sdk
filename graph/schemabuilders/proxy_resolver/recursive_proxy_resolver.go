package proxy_resolver

import (
	"bytes"
	"fmt"
)

type ProxyResolverContext struct {
	name         string
	isRoot       bool
	arguments    map[string]interface{}
	value        map[string]interface{}
	error        error
	responseCb   ProxyResolverResponseCallback
	parent       *ProxyResolverContext
	subtree      []*ProxyResolverContext
	subtreeNames []string
}

type ProxyResolverResponseCallback func(query []byte) map[string]interface{}

var ProxyResolverContextKey = "proxy-resolver-context"

func NewProxyResolverContext(responseCb ProxyResolverResponseCallback) *ProxyResolverContext {
	return &ProxyResolverContext{
		name:      "root",
		isRoot:    true,
		arguments: nil,
		value:     nil,
		error:     nil,
		parent:    nil,
		// root
		responseCb:   responseCb,
		subtree:      make([]*ProxyResolverContext, 0),
		subtreeNames: make([]string, 0),
	}
}

func (prc *ProxyResolverContext) CreateSubtree(typeName string, arguments map[string]interface{}) *ProxyResolverContext {
	sprc := &ProxyResolverContext{
		name:         typeName,
		isRoot:       false,
		arguments:    arguments,
		value:        nil,
		error:        nil,
		responseCb:   nil,
		parent:       prc,
		subtree:      make([]*ProxyResolverContext, 0),
		subtreeNames: make([]string, 0),
	}

	prc.subtree = append(prc.subtree, sprc)
	prc.subtreeNames = append(prc.subtreeNames, typeName)

	return sprc
}

func (prc *ProxyResolverContext) Resolve() (interface{}, error) {
	// run this part only if this PRC is root prc
	if prc.isRoot {
		// if value is already fetched, return cached response
		if prc.value != nil {
			return prc.value, nil
		}

		// else make http request using prc.responseCb(query)
		query := reconstructRootQuery(prc)
		prc.value = prc.responseCb(query.Bytes())

		return prc.value, nil
	}

	// run this part only if this PRC is children prc
	parentValue, err := prc.parent.Resolve()
	if err != nil {
		return nil, err
	}
	asMap, ok := parentValue.(map[string]interface{})
	if !ok {
		panic("parent value is not map[string]interface{}")
	}

	return asMap[prc.name], nil
}

func reconstructRootQuery(prc *ProxyResolverContext) *bytes.Buffer {
	// query
	query := new(bytes.Buffer)

	// write arguments part
	globalArgumentsMap := make(map[string]interface{})

	// turn everything into query
	query.Write([]byte("query"))
	reconstructSubselectionQuery(query, globalArgumentsMap, prc)

	return query
}

func reconstructSubselectionQuery(queryBuf *bytes.Buffer, globalArgumentsMap map[string]interface{}, prc *ProxyResolverContext) {
	// write arguments part
	reconstructSubselectionArgument(queryBuf, globalArgumentsMap, prc.arguments)

	// write subselection parts
	hasSubselection := len(prc.subtree) != 0

	if hasSubselection {
		queryBuf.Write([]byte("{"))
	}
	for subtreeIdx, subtreeItem := range prc.subtree {
		subtreeName := prc.subtreeNames[subtreeIdx]
		queryBuf.Write([]byte(subtreeName))
		reconstructSubselectionQuery(queryBuf, globalArgumentsMap, subtreeItem)
		queryBuf.Write([]byte(","))
	}
	if hasSubselection {
		queryBuf.Write([]byte("}"))
	}
}

func reconstructSubselectionArgument(qb *bytes.Buffer, globalArgumentsMap map[string]interface{}, arguments map[string]interface{}) {
	if arguments == nil || len(arguments) == 0 {
		return
	}

	qb.Write([]byte("("))
	for ak, av := range arguments {
		qb.Write([]byte(ak))
		qb.Write([]byte(":"))
		qb.Write(toCorrectGraphQLArgument(av))
		qb.Write([]byte(","))

		// stash this argument to global arguments map,
		// because all know parameters need to be described in query as well
		globalArgumentsMap[ak] = av
	}
	qb.Write([]byte(")"))
}

func toCorrectGraphQLArgument(argValue interface{}) []byte {
	switch argValue.(type) {
	case string:
		return []byte(fmt.Sprintf("\"%s\"", argValue))
	case int:
		return []byte(fmt.Sprintf("%d", argValue))
	default:
		return []byte(argValue.(string))
	}
}
