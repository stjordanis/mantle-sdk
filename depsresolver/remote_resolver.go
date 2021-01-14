package depsresolver

import (
	"reflect"
)

type RemoteDepsResolver struct {
}

func NewRemoteDepsResolver() DepsResolver {
	return &RemoteDepsResolver{}
}

func (rDepsResolver *RemoteDepsResolver) SetPredefinedState(interface{}) {

}
func (rDepsResolver *RemoteDepsResolver) Emit(interface{}) error {
	return nil
}

func (rDepsResolver *RemoteDepsResolver) GetState() map[string]interface{} {
	return nil
}

func (rDepsResolver *RemoteDepsResolver) Resolve(reflect.Type) interface{} {
	return nil
}

func (rDepsResolver *RemoteDepsResolver) ResolveLatest(reflect.Type) interface{} {
	return nil
}

func (rDepsResolver *RemoteDepsResolver) Dispose() {

}
