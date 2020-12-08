package graph

import (
	"fmt"
	"github.com/terra-project/mantle-sdk/serdes"
	"github.com/terra-project/mantle-sdk/types"
	"reflect"
	"sync"
)

// TODO: optimize this bit, this is mad inefficient
func UnmarshalInternalQueryResult(result *types.GraphQLInternalResult, target interface{}) error {
	targetValue := reflect.Indirect(reflect.ValueOf(target))

	for key, data := range result.Data {
		targetField := targetValue.FieldByName(key)
		if !targetField.IsValid() {
			return fmt.Errorf("invalid target %s", key)
		}

		targetCache := reflect.New(targetField.Type())

		pack, _ := serdes.Serialize(targetField.Type(), data)
		err := serdes.Deserialize(targetField.Type(), pack, targetCache.Interface())

		if err != nil {
			fmt.Println("????", err, targetCache, data)
			continue
		}

		targetField.Set(targetCache.Elem())
	}

	return nil
}

func CreateParallel(len int) *parallelExecutionContext {
	wg := &sync.WaitGroup{}
	wg.Add(len)
	return &parallelExecutionContext{
		RWMutex: sync.RWMutex{},
		idx:     0,
		wg:      wg,
		result:  make([]ParallelExecutionResult, len),
	}
}

type parallelExecutionContext struct {
	sync.RWMutex
	idx int64
	wg *sync.WaitGroup
	result []ParallelExecutionResult
	errorExists bool
	done bool
}

type ParallelExecutionFunc func() (interface{}, error)
type ParallelExecutionResult struct {
	Result interface{}
	Error error
}

func (pec *parallelExecutionContext) Run(f ParallelExecutionFunc) {
	if pec.done {
		panic("cannot add more runners. parallel execution is already done.")
	}

	i := pec.idx
	pec.idx = pec.idx + 1

	// run goroutine
	go func() {
		defer pec.wg.Done()

		r, e := f()

		pec.RWMutex.Lock()
		var result ParallelExecutionResult
		if e != nil {
			result = ParallelExecutionResult{ Error: e }
			pec.errorExists = true
		} else {
			result = ParallelExecutionResult{ Result: r }
		}

		pec.result[i] = result
		pec.RWMutex.Unlock()
	}()
}

func (pec *parallelExecutionContext) HasErrors() bool {
	return pec.errorExists
}

func (pec *parallelExecutionContext) Sync() []ParallelExecutionResult {
	pec.done = true
	pec.wg.Wait()

	return pec.result
}