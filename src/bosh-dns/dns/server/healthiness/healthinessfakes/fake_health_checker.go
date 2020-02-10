// Code generated by counterfeiter. DO NOT EDIT.
package healthinessfakes

import (
	"bosh-dns/dns/server/healthiness"
	"bosh-dns/healthcheck/api"
	"sync"
)

type FakeHealthChecker struct {
	GetStatusStub        func(string) api.HealthResult
	getStatusMutex       sync.RWMutex
	getStatusArgsForCall []struct {
		arg1 string
	}
	getStatusReturns struct {
		result1 api.HealthResult
	}
	getStatusReturnsOnCall map[int]struct {
		result1 api.HealthResult
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthChecker) GetStatus(arg1 string) api.HealthResult {
	fake.getStatusMutex.Lock()
	ret, specificReturn := fake.getStatusReturnsOnCall[len(fake.getStatusArgsForCall)]
	fake.getStatusArgsForCall = append(fake.getStatusArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStatus", []interface{}{arg1})
	fake.getStatusMutex.Unlock()
	if fake.GetStatusStub != nil {
		return fake.GetStatusStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getStatusReturns
	return fakeReturns.result1
}

func (fake *FakeHealthChecker) GetStatusCallCount() int {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return len(fake.getStatusArgsForCall)
}

func (fake *FakeHealthChecker) GetStatusCalls(stub func(string) api.HealthResult) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = stub
}

func (fake *FakeHealthChecker) GetStatusArgsForCall(i int) string {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	argsForCall := fake.getStatusArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHealthChecker) GetStatusReturns(result1 api.HealthResult) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = nil
	fake.getStatusReturns = struct {
		result1 api.HealthResult
	}{result1}
}

func (fake *FakeHealthChecker) GetStatusReturnsOnCall(i int, result1 api.HealthResult) {
	fake.getStatusMutex.Lock()
	defer fake.getStatusMutex.Unlock()
	fake.GetStatusStub = nil
	if fake.getStatusReturnsOnCall == nil {
		fake.getStatusReturnsOnCall = make(map[int]struct {
			result1 api.HealthResult
		})
	}
	fake.getStatusReturnsOnCall[i] = struct {
		result1 api.HealthResult
	}{result1}
}

func (fake *FakeHealthChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHealthChecker) recordInvocation(key string, args []interface{}) {
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

var _ healthiness.HealthChecker = new(FakeHealthChecker)
