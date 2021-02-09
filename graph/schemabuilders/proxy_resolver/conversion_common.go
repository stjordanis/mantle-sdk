package proxy_resolver

import (
	"fmt"
)

type (
	Definitions map[string]TypeDescriptor
)

var (
	errUnknownArgumentType = func(name string) error { return fmt.Errorf("unknown argument type") }
	errUnknownCustomScalar = func(name string) error { return fmt.Errorf("unknown scalar type") }
)
