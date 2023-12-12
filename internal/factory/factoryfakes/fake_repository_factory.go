// Code generated by counterfeiter. DO NOT EDIT.
package factoryfakes

import (
	"sync"

	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

type FakeRepositoryFactory struct {
	FindByDispensaryStub        func(string, string) (repository.Repository, error)
	findByDispensaryMutex       sync.RWMutex
	findByDispensaryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	findByDispensaryReturns struct {
		result1 repository.Repository
		result2 error
	}
	findByDispensaryReturnsOnCall map[int]struct {
		result1 repository.Repository
		result2 error
	}
	FindByDispensaryMenuStub        func(string, string, string) (repository.Repository, error)
	findByDispensaryMenuMutex       sync.RWMutex
	findByDispensaryMenuArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	findByDispensaryMenuReturns struct {
		result1 repository.Repository
		result2 error
	}
	findByDispensaryMenuReturnsOnCall map[int]struct {
		result1 repository.Repository
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepositoryFactory) FindByDispensary(arg1 string, arg2 string) (repository.Repository, error) {
	fake.findByDispensaryMutex.Lock()
	ret, specificReturn := fake.findByDispensaryReturnsOnCall[len(fake.findByDispensaryArgsForCall)]
	fake.findByDispensaryArgsForCall = append(fake.findByDispensaryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.FindByDispensaryStub
	fakeReturns := fake.findByDispensaryReturns
	fake.recordInvocation("FindByDispensary", []interface{}{arg1, arg2})
	fake.findByDispensaryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepositoryFactory) FindByDispensaryCallCount() int {
	fake.findByDispensaryMutex.RLock()
	defer fake.findByDispensaryMutex.RUnlock()
	return len(fake.findByDispensaryArgsForCall)
}

func (fake *FakeRepositoryFactory) FindByDispensaryCalls(stub func(string, string) (repository.Repository, error)) {
	fake.findByDispensaryMutex.Lock()
	defer fake.findByDispensaryMutex.Unlock()
	fake.FindByDispensaryStub = stub
}

func (fake *FakeRepositoryFactory) FindByDispensaryArgsForCall(i int) (string, string) {
	fake.findByDispensaryMutex.RLock()
	defer fake.findByDispensaryMutex.RUnlock()
	argsForCall := fake.findByDispensaryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepositoryFactory) FindByDispensaryReturns(result1 repository.Repository, result2 error) {
	fake.findByDispensaryMutex.Lock()
	defer fake.findByDispensaryMutex.Unlock()
	fake.FindByDispensaryStub = nil
	fake.findByDispensaryReturns = struct {
		result1 repository.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryFactory) FindByDispensaryReturnsOnCall(i int, result1 repository.Repository, result2 error) {
	fake.findByDispensaryMutex.Lock()
	defer fake.findByDispensaryMutex.Unlock()
	fake.FindByDispensaryStub = nil
	if fake.findByDispensaryReturnsOnCall == nil {
		fake.findByDispensaryReturnsOnCall = make(map[int]struct {
			result1 repository.Repository
			result2 error
		})
	}
	fake.findByDispensaryReturnsOnCall[i] = struct {
		result1 repository.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryFactory) FindByDispensaryMenu(arg1 string, arg2 string, arg3 string) (repository.Repository, error) {
	fake.findByDispensaryMenuMutex.Lock()
	ret, specificReturn := fake.findByDispensaryMenuReturnsOnCall[len(fake.findByDispensaryMenuArgsForCall)]
	fake.findByDispensaryMenuArgsForCall = append(fake.findByDispensaryMenuArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FindByDispensaryMenuStub
	fakeReturns := fake.findByDispensaryMenuReturns
	fake.recordInvocation("FindByDispensaryMenu", []interface{}{arg1, arg2, arg3})
	fake.findByDispensaryMenuMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepositoryFactory) FindByDispensaryMenuCallCount() int {
	fake.findByDispensaryMenuMutex.RLock()
	defer fake.findByDispensaryMenuMutex.RUnlock()
	return len(fake.findByDispensaryMenuArgsForCall)
}

func (fake *FakeRepositoryFactory) FindByDispensaryMenuCalls(stub func(string, string, string) (repository.Repository, error)) {
	fake.findByDispensaryMenuMutex.Lock()
	defer fake.findByDispensaryMenuMutex.Unlock()
	fake.FindByDispensaryMenuStub = stub
}

func (fake *FakeRepositoryFactory) FindByDispensaryMenuArgsForCall(i int) (string, string, string) {
	fake.findByDispensaryMenuMutex.RLock()
	defer fake.findByDispensaryMenuMutex.RUnlock()
	argsForCall := fake.findByDispensaryMenuArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepositoryFactory) FindByDispensaryMenuReturns(result1 repository.Repository, result2 error) {
	fake.findByDispensaryMenuMutex.Lock()
	defer fake.findByDispensaryMenuMutex.Unlock()
	fake.FindByDispensaryMenuStub = nil
	fake.findByDispensaryMenuReturns = struct {
		result1 repository.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryFactory) FindByDispensaryMenuReturnsOnCall(i int, result1 repository.Repository, result2 error) {
	fake.findByDispensaryMenuMutex.Lock()
	defer fake.findByDispensaryMenuMutex.Unlock()
	fake.FindByDispensaryMenuStub = nil
	if fake.findByDispensaryMenuReturnsOnCall == nil {
		fake.findByDispensaryMenuReturnsOnCall = make(map[int]struct {
			result1 repository.Repository
			result2 error
		})
	}
	fake.findByDispensaryMenuReturnsOnCall[i] = struct {
		result1 repository.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findByDispensaryMutex.RLock()
	defer fake.findByDispensaryMutex.RUnlock()
	fake.findByDispensaryMenuMutex.RLock()
	defer fake.findByDispensaryMenuMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepositoryFactory) recordInvocation(key string, args []interface{}) {
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

var _ factory.RepositoryFactory = new(FakeRepositoryFactory)
