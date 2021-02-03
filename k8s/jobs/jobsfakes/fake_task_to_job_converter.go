// Code generated by counterfeiter. DO NOT EDIT.
package jobsfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/k8s/jobs"
	"code.cloudfoundry.org/eirini/opi"
	v1 "k8s.io/api/batch/v1"
)

type FakeTaskToJobConverter struct {
	ConvertStub        func(*opi.Task) *v1.Job
	convertMutex       sync.RWMutex
	convertArgsForCall []struct {
		arg1 *opi.Task
	}
	convertReturns struct {
		result1 *v1.Job
	}
	convertReturnsOnCall map[int]struct {
		result1 *v1.Job
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskToJobConverter) Convert(arg1 *opi.Task) *v1.Job {
	fake.convertMutex.Lock()
	ret, specificReturn := fake.convertReturnsOnCall[len(fake.convertArgsForCall)]
	fake.convertArgsForCall = append(fake.convertArgsForCall, struct {
		arg1 *opi.Task
	}{arg1})
	stub := fake.ConvertStub
	fakeReturns := fake.convertReturns
	fake.recordInvocation("Convert", []interface{}{arg1})
	fake.convertMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskToJobConverter) ConvertCallCount() int {
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	return len(fake.convertArgsForCall)
}

func (fake *FakeTaskToJobConverter) ConvertCalls(stub func(*opi.Task) *v1.Job) {
	fake.convertMutex.Lock()
	defer fake.convertMutex.Unlock()
	fake.ConvertStub = stub
}

func (fake *FakeTaskToJobConverter) ConvertArgsForCall(i int) *opi.Task {
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	argsForCall := fake.convertArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskToJobConverter) ConvertReturns(result1 *v1.Job) {
	fake.convertMutex.Lock()
	defer fake.convertMutex.Unlock()
	fake.ConvertStub = nil
	fake.convertReturns = struct {
		result1 *v1.Job
	}{result1}
}

func (fake *FakeTaskToJobConverter) ConvertReturnsOnCall(i int, result1 *v1.Job) {
	fake.convertMutex.Lock()
	defer fake.convertMutex.Unlock()
	fake.ConvertStub = nil
	if fake.convertReturnsOnCall == nil {
		fake.convertReturnsOnCall = make(map[int]struct {
			result1 *v1.Job
		})
	}
	fake.convertReturnsOnCall[i] = struct {
		result1 *v1.Job
	}{result1}
}

func (fake *FakeTaskToJobConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convertMutex.RLock()
	defer fake.convertMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskToJobConverter) recordInvocation(key string, args []interface{}) {
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

var _ jobs.TaskToJobConverter = new(FakeTaskToJobConverter)