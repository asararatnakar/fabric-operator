// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/IBM-Blockchain/fabric-operator/pkg/crd"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Client struct {
	CreateCRDStub        func(*v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error)
	createCRDMutex       sync.RWMutex
	createCRDArgsForCall []struct {
		arg1 *v1.CustomResourceDefinition
	}
	createCRDReturns struct {
		result1 *v1.CustomResourceDefinition
		result2 error
	}
	createCRDReturnsOnCall map[int]struct {
		result1 *v1.CustomResourceDefinition
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Client) CreateCRD(arg1 *v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error) {
	fake.createCRDMutex.Lock()
	ret, specificReturn := fake.createCRDReturnsOnCall[len(fake.createCRDArgsForCall)]
	fake.createCRDArgsForCall = append(fake.createCRDArgsForCall, struct {
		arg1 *v1.CustomResourceDefinition
	}{arg1})
	stub := fake.CreateCRDStub
	fakeReturns := fake.createCRDReturns
	fake.recordInvocation("CreateCRD", []interface{}{arg1})
	fake.createCRDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Client) CreateCRDCallCount() int {
	fake.createCRDMutex.RLock()
	defer fake.createCRDMutex.RUnlock()
	return len(fake.createCRDArgsForCall)
}

func (fake *Client) CreateCRDCalls(stub func(*v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error)) {
	fake.createCRDMutex.Lock()
	defer fake.createCRDMutex.Unlock()
	fake.CreateCRDStub = stub
}

func (fake *Client) CreateCRDArgsForCall(i int) *v1.CustomResourceDefinition {
	fake.createCRDMutex.RLock()
	defer fake.createCRDMutex.RUnlock()
	argsForCall := fake.createCRDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Client) CreateCRDReturns(result1 *v1.CustomResourceDefinition, result2 error) {
	fake.createCRDMutex.Lock()
	defer fake.createCRDMutex.Unlock()
	fake.CreateCRDStub = nil
	fake.createCRDReturns = struct {
		result1 *v1.CustomResourceDefinition
		result2 error
	}{result1, result2}
}

func (fake *Client) CreateCRDReturnsOnCall(i int, result1 *v1.CustomResourceDefinition, result2 error) {
	fake.createCRDMutex.Lock()
	defer fake.createCRDMutex.Unlock()
	fake.CreateCRDStub = nil
	if fake.createCRDReturnsOnCall == nil {
		fake.createCRDReturnsOnCall = make(map[int]struct {
			result1 *v1.CustomResourceDefinition
			result2 error
		})
	}
	fake.createCRDReturnsOnCall[i] = struct {
		result1 *v1.CustomResourceDefinition
		result2 error
	}{result1, result2}
}

func (fake *Client) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCRDMutex.RLock()
	defer fake.createCRDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Client) recordInvocation(key string, args []interface{}) {
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

var _ crd.Client = new(Client)