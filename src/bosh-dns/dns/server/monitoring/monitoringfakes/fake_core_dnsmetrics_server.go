// Code generated by counterfeiter. DO NOT EDIT.
package monitoringfakes

import (
	"bosh-dns/dns/server/monitoring"
	"context"
	"sync"

	"github.com/miekg/dns"
)

type FakeCoreDNSMetricsServer struct {
	OnFinalShutdownStub        func() error
	onFinalShutdownMutex       sync.RWMutex
	onFinalShutdownArgsForCall []struct {
	}
	onFinalShutdownReturns struct {
		result1 error
	}
	onFinalShutdownReturnsOnCall map[int]struct {
		result1 error
	}
	OnStartupStub        func() error
	onStartupMutex       sync.RWMutex
	onStartupArgsForCall []struct {
	}
	onStartupReturns struct {
		result1 error
	}
	onStartupReturnsOnCall map[int]struct {
		result1 error
	}
	ServeDNSStub        func(context.Context, dns.ResponseWriter, *dns.Msg) (int, error)
	serveDNSMutex       sync.RWMutex
	serveDNSArgsForCall []struct {
		arg1 context.Context
		arg2 dns.ResponseWriter
		arg3 *dns.Msg
	}
	serveDNSReturns struct {
		result1 int
		result2 error
	}
	serveDNSReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCoreDNSMetricsServer) OnFinalShutdown() error {
	fake.onFinalShutdownMutex.Lock()
	ret, specificReturn := fake.onFinalShutdownReturnsOnCall[len(fake.onFinalShutdownArgsForCall)]
	fake.onFinalShutdownArgsForCall = append(fake.onFinalShutdownArgsForCall, struct {
	}{})
	fake.recordInvocation("OnFinalShutdown", []interface{}{})
	fake.onFinalShutdownMutex.Unlock()
	if fake.OnFinalShutdownStub != nil {
		return fake.OnFinalShutdownStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.onFinalShutdownReturns
	return fakeReturns.result1
}

func (fake *FakeCoreDNSMetricsServer) OnFinalShutdownCallCount() int {
	fake.onFinalShutdownMutex.RLock()
	defer fake.onFinalShutdownMutex.RUnlock()
	return len(fake.onFinalShutdownArgsForCall)
}

func (fake *FakeCoreDNSMetricsServer) OnFinalShutdownCalls(stub func() error) {
	fake.onFinalShutdownMutex.Lock()
	defer fake.onFinalShutdownMutex.Unlock()
	fake.OnFinalShutdownStub = stub
}

func (fake *FakeCoreDNSMetricsServer) OnFinalShutdownReturns(result1 error) {
	fake.onFinalShutdownMutex.Lock()
	defer fake.onFinalShutdownMutex.Unlock()
	fake.OnFinalShutdownStub = nil
	fake.onFinalShutdownReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCoreDNSMetricsServer) OnFinalShutdownReturnsOnCall(i int, result1 error) {
	fake.onFinalShutdownMutex.Lock()
	defer fake.onFinalShutdownMutex.Unlock()
	fake.OnFinalShutdownStub = nil
	if fake.onFinalShutdownReturnsOnCall == nil {
		fake.onFinalShutdownReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.onFinalShutdownReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCoreDNSMetricsServer) OnStartup() error {
	fake.onStartupMutex.Lock()
	ret, specificReturn := fake.onStartupReturnsOnCall[len(fake.onStartupArgsForCall)]
	fake.onStartupArgsForCall = append(fake.onStartupArgsForCall, struct {
	}{})
	fake.recordInvocation("OnStartup", []interface{}{})
	fake.onStartupMutex.Unlock()
	if fake.OnStartupStub != nil {
		return fake.OnStartupStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.onStartupReturns
	return fakeReturns.result1
}

func (fake *FakeCoreDNSMetricsServer) OnStartupCallCount() int {
	fake.onStartupMutex.RLock()
	defer fake.onStartupMutex.RUnlock()
	return len(fake.onStartupArgsForCall)
}

func (fake *FakeCoreDNSMetricsServer) OnStartupCalls(stub func() error) {
	fake.onStartupMutex.Lock()
	defer fake.onStartupMutex.Unlock()
	fake.OnStartupStub = stub
}

func (fake *FakeCoreDNSMetricsServer) OnStartupReturns(result1 error) {
	fake.onStartupMutex.Lock()
	defer fake.onStartupMutex.Unlock()
	fake.OnStartupStub = nil
	fake.onStartupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCoreDNSMetricsServer) OnStartupReturnsOnCall(i int, result1 error) {
	fake.onStartupMutex.Lock()
	defer fake.onStartupMutex.Unlock()
	fake.OnStartupStub = nil
	if fake.onStartupReturnsOnCall == nil {
		fake.onStartupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.onStartupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCoreDNSMetricsServer) ServeDNS(arg1 context.Context, arg2 dns.ResponseWriter, arg3 *dns.Msg) (int, error) {
	fake.serveDNSMutex.Lock()
	ret, specificReturn := fake.serveDNSReturnsOnCall[len(fake.serveDNSArgsForCall)]
	fake.serveDNSArgsForCall = append(fake.serveDNSArgsForCall, struct {
		arg1 context.Context
		arg2 dns.ResponseWriter
		arg3 *dns.Msg
	}{arg1, arg2, arg3})
	fake.recordInvocation("ServeDNS", []interface{}{arg1, arg2, arg3})
	fake.serveDNSMutex.Unlock()
	if fake.ServeDNSStub != nil {
		return fake.ServeDNSStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.serveDNSReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCoreDNSMetricsServer) ServeDNSCallCount() int {
	fake.serveDNSMutex.RLock()
	defer fake.serveDNSMutex.RUnlock()
	return len(fake.serveDNSArgsForCall)
}

func (fake *FakeCoreDNSMetricsServer) ServeDNSCalls(stub func(context.Context, dns.ResponseWriter, *dns.Msg) (int, error)) {
	fake.serveDNSMutex.Lock()
	defer fake.serveDNSMutex.Unlock()
	fake.ServeDNSStub = stub
}

func (fake *FakeCoreDNSMetricsServer) ServeDNSArgsForCall(i int) (context.Context, dns.ResponseWriter, *dns.Msg) {
	fake.serveDNSMutex.RLock()
	defer fake.serveDNSMutex.RUnlock()
	argsForCall := fake.serveDNSArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCoreDNSMetricsServer) ServeDNSReturns(result1 int, result2 error) {
	fake.serveDNSMutex.Lock()
	defer fake.serveDNSMutex.Unlock()
	fake.ServeDNSStub = nil
	fake.serveDNSReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreDNSMetricsServer) ServeDNSReturnsOnCall(i int, result1 int, result2 error) {
	fake.serveDNSMutex.Lock()
	defer fake.serveDNSMutex.Unlock()
	fake.ServeDNSStub = nil
	if fake.serveDNSReturnsOnCall == nil {
		fake.serveDNSReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.serveDNSReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeCoreDNSMetricsServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onFinalShutdownMutex.RLock()
	defer fake.onFinalShutdownMutex.RUnlock()
	fake.onStartupMutex.RLock()
	defer fake.onStartupMutex.RUnlock()
	fake.serveDNSMutex.RLock()
	defer fake.serveDNSMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCoreDNSMetricsServer) recordInvocation(key string, args []interface{}) {
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

var _ monitoring.CoreDNSMetricsServer = new(FakeCoreDNSMetricsServer)
