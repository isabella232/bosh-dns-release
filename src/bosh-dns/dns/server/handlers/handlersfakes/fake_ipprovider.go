// Code generated by counterfeiter. DO NOT EDIT.
package handlersfakes

import (
	"bosh-dns/dns/server/handlers"
	"sync"
)

type FakeIPProvider struct {
	AllIPsStub        func() map[string]bool
	allIPsMutex       sync.RWMutex
	allIPsArgsForCall []struct{}
	allIPsReturns     struct {
		result1 map[string]bool
	}
	allIPsReturnsOnCall map[int]struct {
		result1 map[string]bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIPProvider) AllIPs() map[string]bool {
	fake.allIPsMutex.Lock()
	ret, specificReturn := fake.allIPsReturnsOnCall[len(fake.allIPsArgsForCall)]
	fake.allIPsArgsForCall = append(fake.allIPsArgsForCall, struct{}{})
	fake.recordInvocation("AllIPs", []interface{}{})
	fake.allIPsMutex.Unlock()
	if fake.AllIPsStub != nil {
		return fake.AllIPsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.allIPsReturns.result1
}

func (fake *FakeIPProvider) AllIPsCallCount() int {
	fake.allIPsMutex.RLock()
	defer fake.allIPsMutex.RUnlock()
	return len(fake.allIPsArgsForCall)
}

func (fake *FakeIPProvider) AllIPsReturns(result1 map[string]bool) {
	fake.AllIPsStub = nil
	fake.allIPsReturns = struct {
		result1 map[string]bool
	}{result1}
}

func (fake *FakeIPProvider) AllIPsReturnsOnCall(i int, result1 map[string]bool) {
	fake.AllIPsStub = nil
	if fake.allIPsReturnsOnCall == nil {
		fake.allIPsReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
		})
	}
	fake.allIPsReturnsOnCall[i] = struct {
		result1 map[string]bool
	}{result1}
}

func (fake *FakeIPProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allIPsMutex.RLock()
	defer fake.allIPsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIPProvider) recordInvocation(key string, args []interface{}) {
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

var _ handlers.IPProvider = new(FakeIPProvider)