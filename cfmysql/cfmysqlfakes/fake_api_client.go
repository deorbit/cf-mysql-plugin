// Code generated by counterfeiter. DO NOT EDIT.
package cfmysqlfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
	sdkModels "code.cloudfoundry.org/cli/plugin/models"
	"github.com/deorbit/cf-mysql-plugin/cfmysql"
	pluginModels "github.com/deorbit/cf-mysql-plugin/cfmysql/models"
)

type FakeApiClient struct {
	GetStartedAppsStub        func(cliConnection plugin.CliConnection) ([]sdkModels.GetAppsModel, error)
	getStartedAppsMutex       sync.RWMutex
	getStartedAppsArgsForCall []struct {
		cliConnection plugin.CliConnection
	}
	getStartedAppsReturns struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}
	getStartedAppsReturnsOnCall map[int]struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}
	GetServiceStub        func(cliConnection plugin.CliConnection, spaceGuid string, name string) (pluginModels.ServiceInstance, error)
	getServiceMutex       sync.RWMutex
	getServiceArgsForCall []struct {
		cliConnection plugin.CliConnection
		spaceGuid     string
		name          string
	}
	getServiceReturns struct {
		result1 pluginModels.ServiceInstance
		result2 error
	}
	getServiceReturnsOnCall map[int]struct {
		result1 pluginModels.ServiceInstance
		result2 error
	}
	GetServiceKeyStub        func(cliConnection plugin.CliConnection, serviceInstanceGuid string, keyName string) (key pluginModels.ServiceKey, found bool, err error)
	getServiceKeyMutex       sync.RWMutex
	getServiceKeyArgsForCall []struct {
		cliConnection       plugin.CliConnection
		serviceInstanceGuid string
		keyName             string
	}
	getServiceKeyReturns struct {
		result1 pluginModels.ServiceKey
		result2 bool
		result3 error
	}
	getServiceKeyReturnsOnCall map[int]struct {
		result1 pluginModels.ServiceKey
		result2 bool
		result3 error
	}
	CreateServiceKeyStub        func(cliConnection plugin.CliConnection, serviceInstanceGuid string, keyName string) (pluginModels.ServiceKey, error)
	createServiceKeyMutex       sync.RWMutex
	createServiceKeyArgsForCall []struct {
		cliConnection       plugin.CliConnection
		serviceInstanceGuid string
		keyName             string
	}
	createServiceKeyReturns struct {
		result1 pluginModels.ServiceKey
		result2 error
	}
	createServiceKeyReturnsOnCall map[int]struct {
		result1 pluginModels.ServiceKey
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApiClient) GetStartedApps(cliConnection plugin.CliConnection) ([]sdkModels.GetAppsModel, error) {
	fake.getStartedAppsMutex.Lock()
	ret, specificReturn := fake.getStartedAppsReturnsOnCall[len(fake.getStartedAppsArgsForCall)]
	fake.getStartedAppsArgsForCall = append(fake.getStartedAppsArgsForCall, struct {
		cliConnection plugin.CliConnection
	}{cliConnection})
	fake.recordInvocation("GetStartedApps", []interface{}{cliConnection})
	fake.getStartedAppsMutex.Unlock()
	if fake.GetStartedAppsStub != nil {
		return fake.GetStartedAppsStub(cliConnection)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStartedAppsReturns.result1, fake.getStartedAppsReturns.result2
}

func (fake *FakeApiClient) GetStartedAppsCallCount() int {
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	return len(fake.getStartedAppsArgsForCall)
}

func (fake *FakeApiClient) GetStartedAppsArgsForCall(i int) plugin.CliConnection {
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	return fake.getStartedAppsArgsForCall[i].cliConnection
}

func (fake *FakeApiClient) GetStartedAppsReturns(result1 []sdkModels.GetAppsModel, result2 error) {
	fake.GetStartedAppsStub = nil
	fake.getStartedAppsReturns = struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetStartedAppsReturnsOnCall(i int, result1 []sdkModels.GetAppsModel, result2 error) {
	fake.GetStartedAppsStub = nil
	if fake.getStartedAppsReturnsOnCall == nil {
		fake.getStartedAppsReturnsOnCall = make(map[int]struct {
			result1 []sdkModels.GetAppsModel
			result2 error
		})
	}
	fake.getStartedAppsReturnsOnCall[i] = struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetService(cliConnection plugin.CliConnection, spaceGuid string, name string) (pluginModels.ServiceInstance, error) {
	fake.getServiceMutex.Lock()
	ret, specificReturn := fake.getServiceReturnsOnCall[len(fake.getServiceArgsForCall)]
	fake.getServiceArgsForCall = append(fake.getServiceArgsForCall, struct {
		cliConnection plugin.CliConnection
		spaceGuid     string
		name          string
	}{cliConnection, spaceGuid, name})
	fake.recordInvocation("GetService", []interface{}{cliConnection, spaceGuid, name})
	fake.getServiceMutex.Unlock()
	if fake.GetServiceStub != nil {
		return fake.GetServiceStub(cliConnection, spaceGuid, name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceReturns.result1, fake.getServiceReturns.result2
}

func (fake *FakeApiClient) GetServiceCallCount() int {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return len(fake.getServiceArgsForCall)
}

func (fake *FakeApiClient) GetServiceArgsForCall(i int) (plugin.CliConnection, string, string) {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return fake.getServiceArgsForCall[i].cliConnection, fake.getServiceArgsForCall[i].spaceGuid, fake.getServiceArgsForCall[i].name
}

func (fake *FakeApiClient) GetServiceReturns(result1 pluginModels.ServiceInstance, result2 error) {
	fake.GetServiceStub = nil
	fake.getServiceReturns = struct {
		result1 pluginModels.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetServiceReturnsOnCall(i int, result1 pluginModels.ServiceInstance, result2 error) {
	fake.GetServiceStub = nil
	if fake.getServiceReturnsOnCall == nil {
		fake.getServiceReturnsOnCall = make(map[int]struct {
			result1 pluginModels.ServiceInstance
			result2 error
		})
	}
	fake.getServiceReturnsOnCall[i] = struct {
		result1 pluginModels.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetServiceKey(cliConnection plugin.CliConnection, serviceInstanceGuid string, keyName string) (key pluginModels.ServiceKey, found bool, err error) {
	fake.getServiceKeyMutex.Lock()
	ret, specificReturn := fake.getServiceKeyReturnsOnCall[len(fake.getServiceKeyArgsForCall)]
	fake.getServiceKeyArgsForCall = append(fake.getServiceKeyArgsForCall, struct {
		cliConnection       plugin.CliConnection
		serviceInstanceGuid string
		keyName             string
	}{cliConnection, serviceInstanceGuid, keyName})
	fake.recordInvocation("GetServiceKey", []interface{}{cliConnection, serviceInstanceGuid, keyName})
	fake.getServiceKeyMutex.Unlock()
	if fake.GetServiceKeyStub != nil {
		return fake.GetServiceKeyStub(cliConnection, serviceInstanceGuid, keyName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceKeyReturns.result1, fake.getServiceKeyReturns.result2, fake.getServiceKeyReturns.result3
}

func (fake *FakeApiClient) GetServiceKeyCallCount() int {
	fake.getServiceKeyMutex.RLock()
	defer fake.getServiceKeyMutex.RUnlock()
	return len(fake.getServiceKeyArgsForCall)
}

func (fake *FakeApiClient) GetServiceKeyArgsForCall(i int) (plugin.CliConnection, string, string) {
	fake.getServiceKeyMutex.RLock()
	defer fake.getServiceKeyMutex.RUnlock()
	return fake.getServiceKeyArgsForCall[i].cliConnection, fake.getServiceKeyArgsForCall[i].serviceInstanceGuid, fake.getServiceKeyArgsForCall[i].keyName
}

func (fake *FakeApiClient) GetServiceKeyReturns(result1 pluginModels.ServiceKey, result2 bool, result3 error) {
	fake.GetServiceKeyStub = nil
	fake.getServiceKeyReturns = struct {
		result1 pluginModels.ServiceKey
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeApiClient) GetServiceKeyReturnsOnCall(i int, result1 pluginModels.ServiceKey, result2 bool, result3 error) {
	fake.GetServiceKeyStub = nil
	if fake.getServiceKeyReturnsOnCall == nil {
		fake.getServiceKeyReturnsOnCall = make(map[int]struct {
			result1 pluginModels.ServiceKey
			result2 bool
			result3 error
		})
	}
	fake.getServiceKeyReturnsOnCall[i] = struct {
		result1 pluginModels.ServiceKey
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeApiClient) CreateServiceKey(cliConnection plugin.CliConnection, serviceInstanceGuid string, keyName string) (pluginModels.ServiceKey, error) {
	fake.createServiceKeyMutex.Lock()
	ret, specificReturn := fake.createServiceKeyReturnsOnCall[len(fake.createServiceKeyArgsForCall)]
	fake.createServiceKeyArgsForCall = append(fake.createServiceKeyArgsForCall, struct {
		cliConnection       plugin.CliConnection
		serviceInstanceGuid string
		keyName             string
	}{cliConnection, serviceInstanceGuid, keyName})
	fake.recordInvocation("CreateServiceKey", []interface{}{cliConnection, serviceInstanceGuid, keyName})
	fake.createServiceKeyMutex.Unlock()
	if fake.CreateServiceKeyStub != nil {
		return fake.CreateServiceKeyStub(cliConnection, serviceInstanceGuid, keyName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createServiceKeyReturns.result1, fake.createServiceKeyReturns.result2
}

func (fake *FakeApiClient) CreateServiceKeyCallCount() int {
	fake.createServiceKeyMutex.RLock()
	defer fake.createServiceKeyMutex.RUnlock()
	return len(fake.createServiceKeyArgsForCall)
}

func (fake *FakeApiClient) CreateServiceKeyArgsForCall(i int) (plugin.CliConnection, string, string) {
	fake.createServiceKeyMutex.RLock()
	defer fake.createServiceKeyMutex.RUnlock()
	return fake.createServiceKeyArgsForCall[i].cliConnection, fake.createServiceKeyArgsForCall[i].serviceInstanceGuid, fake.createServiceKeyArgsForCall[i].keyName
}

func (fake *FakeApiClient) CreateServiceKeyReturns(result1 pluginModels.ServiceKey, result2 error) {
	fake.CreateServiceKeyStub = nil
	fake.createServiceKeyReturns = struct {
		result1 pluginModels.ServiceKey
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) CreateServiceKeyReturnsOnCall(i int, result1 pluginModels.ServiceKey, result2 error) {
	fake.CreateServiceKeyStub = nil
	if fake.createServiceKeyReturnsOnCall == nil {
		fake.createServiceKeyReturnsOnCall = make(map[int]struct {
			result1 pluginModels.ServiceKey
			result2 error
		})
	}
	fake.createServiceKeyReturnsOnCall[i] = struct {
		result1 pluginModels.ServiceKey
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	fake.getServiceKeyMutex.RLock()
	defer fake.getServiceKeyMutex.RUnlock()
	fake.createServiceKeyMutex.RLock()
	defer fake.createServiceKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApiClient) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.ApiClient = new(FakeApiClient)
