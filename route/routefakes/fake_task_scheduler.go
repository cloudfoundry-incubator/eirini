// Code generated by counterfeiter. DO NOT EDIT.
package routefakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/eirini/route"
)

type FakeTaskScheduler struct {
	ScheduleStub        func(task func() error)
	scheduleMutex       sync.RWMutex
	scheduleArgsForCall []struct {
		task func() error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskScheduler) Schedule(task func() error) {
	fake.scheduleMutex.Lock()
	fake.scheduleArgsForCall = append(fake.scheduleArgsForCall, struct {
		task func() error
	}{task})
	fake.recordInvocation("Schedule", []interface{}{task})
	fake.scheduleMutex.Unlock()
	if fake.ScheduleStub != nil {
		fake.ScheduleStub(task)
	}
}

func (fake *FakeTaskScheduler) ScheduleCallCount() int {
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	return len(fake.scheduleArgsForCall)
}

func (fake *FakeTaskScheduler) ScheduleArgsForCall(i int) func() error {
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	return fake.scheduleArgsForCall[i].task
}

func (fake *FakeTaskScheduler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskScheduler) recordInvocation(key string, args []interface{}) {
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

var _ route.TaskScheduler = new(FakeTaskScheduler)
