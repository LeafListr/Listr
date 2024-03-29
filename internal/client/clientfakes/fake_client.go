// Code generated by counterfeiter. DO NOT EDIT.
package clientfakes

import (
	"context"
	"sync"

	"github.com/Linkinlog/LeafListr/internal/client"
)

type FakeClient struct {
	QueryStub        func(context.Context, string, string) ([]byte, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	queryReturns struct {
		result1 []byte
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SetEndpointStub        func(client.Endpoint)
	setEndpointMutex       sync.RWMutex
	setEndpointArgsForCall []struct {
		arg1 client.Endpoint
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Query(arg1 context.Context, arg2 string, arg3 string) ([]byte, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.QueryStub
	fakeReturns := fake.queryReturns
	fake.recordInvocation("Query", []interface{}{arg1, arg2, arg3})
	fake.queryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeClient) QueryCalls(stub func(context.Context, string, string) ([]byte, error)) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = stub
}

func (fake *FakeClient) QueryArgsForCall(i int) (context.Context, string, string) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	argsForCall := fake.queryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) QueryReturns(result1 []byte, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) QueryReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SetEndpoint(arg1 client.Endpoint) {
	fake.setEndpointMutex.Lock()
	fake.setEndpointArgsForCall = append(fake.setEndpointArgsForCall, struct {
		arg1 client.Endpoint
	}{arg1})
	stub := fake.SetEndpointStub
	fake.recordInvocation("SetEndpoint", []interface{}{arg1})
	fake.setEndpointMutex.Unlock()
	if stub != nil {
		fake.SetEndpointStub(arg1)
	}
}

func (fake *FakeClient) SetEndpointCallCount() int {
	fake.setEndpointMutex.RLock()
	defer fake.setEndpointMutex.RUnlock()
	return len(fake.setEndpointArgsForCall)
}

func (fake *FakeClient) SetEndpointCalls(stub func(client.Endpoint)) {
	fake.setEndpointMutex.Lock()
	defer fake.setEndpointMutex.Unlock()
	fake.SetEndpointStub = stub
}

func (fake *FakeClient) SetEndpointArgsForCall(i int) client.Endpoint {
	fake.setEndpointMutex.RLock()
	defer fake.setEndpointMutex.RUnlock()
	argsForCall := fake.setEndpointArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.setEndpointMutex.RLock()
	defer fake.setEndpointMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ client.Client = new(FakeClient)
