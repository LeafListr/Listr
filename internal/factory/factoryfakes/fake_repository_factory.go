// Code generated by counterfeiter. DO NOT EDIT.
package factoryfakes

import (
	"sync"

	"github.com/Linkinlog/LeafListr/internal/factory"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

type FakeRepositoryFactory struct {
	FindByDispensaryStub        func() (repository.Repository, error)
	findByDispensaryMutex       sync.RWMutex
	findByDispensaryArgsForCall []struct {
	}
	findByDispensaryReturns struct {
		result1 repository.Repository
		result2 error
	}
	findByDispensaryReturnsOnCall map[int]struct {
		result1 repository.Repository
		result2 error
	}
	FindByDispensaryMenuStub        func() (repository.Repository, error)
	findByDispensaryMenuMutex       sync.RWMutex
	findByDispensaryMenuArgsForCall []struct {
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

func (fake *FakeRepositoryFactory) FindByDispensary() (repository.Repository, error) {
	fake.findByDispensaryMutex.Lock()
	ret, specificReturn := fake.findByDispensaryReturnsOnCall[len(fake.findByDispensaryArgsForCall)]
	fake.findByDispensaryArgsForCall = append(fake.findByDispensaryArgsForCall, struct {
	}{})
	stub := fake.FindByDispensaryStub
	fakeReturns := fake.findByDispensaryReturns
	fake.recordInvocation("FindByDispensary", []interface{}{})
	fake.findByDispensaryMutex.Unlock()
	if stub != nil {
		return stub()
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

func (fake *FakeRepositoryFactory) FindByDispensaryCalls(stub func() (repository.Repository, error)) {
	fake.findByDispensaryMutex.Lock()
	defer fake.findByDispensaryMutex.Unlock()
	fake.FindByDispensaryStub = stub
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

func (fake *FakeRepositoryFactory) FindByDispensaryMenu() (repository.Repository, error) {
	fake.findByDispensaryMenuMutex.Lock()
	ret, specificReturn := fake.findByDispensaryMenuReturnsOnCall[len(fake.findByDispensaryMenuArgsForCall)]
	fake.findByDispensaryMenuArgsForCall = append(fake.findByDispensaryMenuArgsForCall, struct {
	}{})
	stub := fake.FindByDispensaryMenuStub
	fakeReturns := fake.findByDispensaryMenuReturns
	fake.recordInvocation("FindByDispensaryMenu", []interface{}{})
	fake.findByDispensaryMenuMutex.Unlock()
	if stub != nil {
		return stub()
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

func (fake *FakeRepositoryFactory) FindByDispensaryMenuCalls(stub func() (repository.Repository, error)) {
	fake.findByDispensaryMenuMutex.Lock()
	defer fake.findByDispensaryMenuMutex.Unlock()
	fake.FindByDispensaryMenuStub = stub
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
