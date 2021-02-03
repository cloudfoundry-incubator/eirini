// Code generated by counterfeiter. DO NOT EDIT.
package bifrostfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/bifrost"
	"code.cloudfoundry.org/eirini/k8s/shared"
	"code.cloudfoundry.org/eirini/opi"
)

type FakeLRPClient struct {
	DesireStub        func(string, *opi.LRP, ...shared.Option) error
	desireMutex       sync.RWMutex
	desireArgsForCall []struct {
		arg1 string
		arg2 *opi.LRP
		arg3 []shared.Option
	}
	desireReturns struct {
		result1 error
	}
	desireReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(opi.LRPIdentifier) (*opi.LRP, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 opi.LRPIdentifier
	}
	getReturns struct {
		result1 *opi.LRP
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *opi.LRP
		result2 error
	}
	GetInstancesStub        func(opi.LRPIdentifier) ([]*opi.Instance, error)
	getInstancesMutex       sync.RWMutex
	getInstancesArgsForCall []struct {
		arg1 opi.LRPIdentifier
	}
	getInstancesReturns struct {
		result1 []*opi.Instance
		result2 error
	}
	getInstancesReturnsOnCall map[int]struct {
		result1 []*opi.Instance
		result2 error
	}
	ListStub        func() ([]*opi.LRP, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 []*opi.LRP
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []*opi.LRP
		result2 error
	}
	StopStub        func(opi.LRPIdentifier) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		arg1 opi.LRPIdentifier
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	StopInstanceStub        func(opi.LRPIdentifier, uint) error
	stopInstanceMutex       sync.RWMutex
	stopInstanceArgsForCall []struct {
		arg1 opi.LRPIdentifier
		arg2 uint
	}
	stopInstanceReturns struct {
		result1 error
	}
	stopInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(*opi.LRP) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 *opi.LRP
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLRPClient) Desire(arg1 string, arg2 *opi.LRP, arg3 ...shared.Option) error {
	fake.desireMutex.Lock()
	ret, specificReturn := fake.desireReturnsOnCall[len(fake.desireArgsForCall)]
	fake.desireArgsForCall = append(fake.desireArgsForCall, struct {
		arg1 string
		arg2 *opi.LRP
		arg3 []shared.Option
	}{arg1, arg2, arg3})
	stub := fake.DesireStub
	fakeReturns := fake.desireReturns
	fake.recordInvocation("Desire", []interface{}{arg1, arg2, arg3})
	fake.desireMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPClient) DesireCallCount() int {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	return len(fake.desireArgsForCall)
}

func (fake *FakeLRPClient) DesireCalls(stub func(string, *opi.LRP, ...shared.Option) error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = stub
}

func (fake *FakeLRPClient) DesireArgsForCall(i int) (string, *opi.LRP, []shared.Option) {
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	argsForCall := fake.desireArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLRPClient) DesireReturns(result1 error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = nil
	fake.desireReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) DesireReturnsOnCall(i int, result1 error) {
	fake.desireMutex.Lock()
	defer fake.desireMutex.Unlock()
	fake.DesireStub = nil
	if fake.desireReturnsOnCall == nil {
		fake.desireReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.desireReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) Get(arg1 opi.LRPIdentifier) (*opi.LRP, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 opi.LRPIdentifier
	}{arg1})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLRPClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeLRPClient) GetCalls(stub func(opi.LRPIdentifier) (*opi.LRP, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeLRPClient) GetArgsForCall(i int) opi.LRPIdentifier {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLRPClient) GetReturns(result1 *opi.LRP, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *opi.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPClient) GetReturnsOnCall(i int, result1 *opi.LRP, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *opi.LRP
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *opi.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPClient) GetInstances(arg1 opi.LRPIdentifier) ([]*opi.Instance, error) {
	fake.getInstancesMutex.Lock()
	ret, specificReturn := fake.getInstancesReturnsOnCall[len(fake.getInstancesArgsForCall)]
	fake.getInstancesArgsForCall = append(fake.getInstancesArgsForCall, struct {
		arg1 opi.LRPIdentifier
	}{arg1})
	stub := fake.GetInstancesStub
	fakeReturns := fake.getInstancesReturns
	fake.recordInvocation("GetInstances", []interface{}{arg1})
	fake.getInstancesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLRPClient) GetInstancesCallCount() int {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	return len(fake.getInstancesArgsForCall)
}

func (fake *FakeLRPClient) GetInstancesCalls(stub func(opi.LRPIdentifier) ([]*opi.Instance, error)) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = stub
}

func (fake *FakeLRPClient) GetInstancesArgsForCall(i int) opi.LRPIdentifier {
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	argsForCall := fake.getInstancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLRPClient) GetInstancesReturns(result1 []*opi.Instance, result2 error) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = nil
	fake.getInstancesReturns = struct {
		result1 []*opi.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPClient) GetInstancesReturnsOnCall(i int, result1 []*opi.Instance, result2 error) {
	fake.getInstancesMutex.Lock()
	defer fake.getInstancesMutex.Unlock()
	fake.GetInstancesStub = nil
	if fake.getInstancesReturnsOnCall == nil {
		fake.getInstancesReturnsOnCall = make(map[int]struct {
			result1 []*opi.Instance
			result2 error
		})
	}
	fake.getInstancesReturnsOnCall[i] = struct {
		result1 []*opi.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPClient) List() ([]*opi.LRP, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLRPClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeLRPClient) ListCalls(stub func() ([]*opi.LRP, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeLRPClient) ListReturns(result1 []*opi.LRP, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []*opi.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPClient) ListReturnsOnCall(i int, result1 []*opi.LRP, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []*opi.LRP
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []*opi.LRP
		result2 error
	}{result1, result2}
}

func (fake *FakeLRPClient) Stop(arg1 opi.LRPIdentifier) error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		arg1 opi.LRPIdentifier
	}{arg1})
	stub := fake.StopStub
	fakeReturns := fake.stopReturns
	fake.recordInvocation("Stop", []interface{}{arg1})
	fake.stopMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeLRPClient) StopCalls(stub func(opi.LRPIdentifier) error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeLRPClient) StopArgsForCall(i int) opi.LRPIdentifier {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	argsForCall := fake.stopArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLRPClient) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) StopInstance(arg1 opi.LRPIdentifier, arg2 uint) error {
	fake.stopInstanceMutex.Lock()
	ret, specificReturn := fake.stopInstanceReturnsOnCall[len(fake.stopInstanceArgsForCall)]
	fake.stopInstanceArgsForCall = append(fake.stopInstanceArgsForCall, struct {
		arg1 opi.LRPIdentifier
		arg2 uint
	}{arg1, arg2})
	stub := fake.StopInstanceStub
	fakeReturns := fake.stopInstanceReturns
	fake.recordInvocation("StopInstance", []interface{}{arg1, arg2})
	fake.stopInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPClient) StopInstanceCallCount() int {
	fake.stopInstanceMutex.RLock()
	defer fake.stopInstanceMutex.RUnlock()
	return len(fake.stopInstanceArgsForCall)
}

func (fake *FakeLRPClient) StopInstanceCalls(stub func(opi.LRPIdentifier, uint) error) {
	fake.stopInstanceMutex.Lock()
	defer fake.stopInstanceMutex.Unlock()
	fake.StopInstanceStub = stub
}

func (fake *FakeLRPClient) StopInstanceArgsForCall(i int) (opi.LRPIdentifier, uint) {
	fake.stopInstanceMutex.RLock()
	defer fake.stopInstanceMutex.RUnlock()
	argsForCall := fake.stopInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLRPClient) StopInstanceReturns(result1 error) {
	fake.stopInstanceMutex.Lock()
	defer fake.stopInstanceMutex.Unlock()
	fake.StopInstanceStub = nil
	fake.stopInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) StopInstanceReturnsOnCall(i int, result1 error) {
	fake.stopInstanceMutex.Lock()
	defer fake.stopInstanceMutex.Unlock()
	fake.StopInstanceStub = nil
	if fake.stopInstanceReturnsOnCall == nil {
		fake.stopInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) Update(arg1 *opi.LRP) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 *opi.LRP
	}{arg1})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLRPClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeLRPClient) UpdateCalls(stub func(*opi.LRP) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeLRPClient) UpdateArgsForCall(i int) *opi.LRP {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLRPClient) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLRPClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.desireMutex.RLock()
	defer fake.desireMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getInstancesMutex.RLock()
	defer fake.getInstancesMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.stopInstanceMutex.RLock()
	defer fake.stopInstanceMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLRPClient) recordInvocation(key string, args []interface{}) {
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

var _ bifrost.LRPClient = new(FakeLRPClient)