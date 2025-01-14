// Code generated by counterfeiter. DO NOT EDIT.
package cfmysqlfakes

import (
	"os"
	"sync"

	"github.com/deorbit/cf-mysql-plugin/cfmysql"
)

type FakeIoUtilWrapper struct {
	TempFileStub        func(dir, pattern string) (f *os.File, err error)
	tempFileMutex       sync.RWMutex
	tempFileArgsForCall []struct {
		dir     string
		pattern string
	}
	tempFileReturns struct {
		result1 *os.File
		result2 error
	}
	tempFileReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIoUtilWrapper) TempFile(dir string, pattern string) (f *os.File, err error) {
	fake.tempFileMutex.Lock()
	ret, specificReturn := fake.tempFileReturnsOnCall[len(fake.tempFileArgsForCall)]
	fake.tempFileArgsForCall = append(fake.tempFileArgsForCall, struct {
		dir     string
		pattern string
	}{dir, pattern})
	fake.recordInvocation("TempFile", []interface{}{dir, pattern})
	fake.tempFileMutex.Unlock()
	if fake.TempFileStub != nil {
		return fake.TempFileStub(dir, pattern)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.tempFileReturns.result1, fake.tempFileReturns.result2
}

func (fake *FakeIoUtilWrapper) TempFileCallCount() int {
	fake.tempFileMutex.RLock()
	defer fake.tempFileMutex.RUnlock()
	return len(fake.tempFileArgsForCall)
}

func (fake *FakeIoUtilWrapper) TempFileArgsForCall(i int) (string, string) {
	fake.tempFileMutex.RLock()
	defer fake.tempFileMutex.RUnlock()
	return fake.tempFileArgsForCall[i].dir, fake.tempFileArgsForCall[i].pattern
}

func (fake *FakeIoUtilWrapper) TempFileReturns(result1 *os.File, result2 error) {
	fake.TempFileStub = nil
	fake.tempFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeIoUtilWrapper) TempFileReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.TempFileStub = nil
	if fake.tempFileReturnsOnCall == nil {
		fake.tempFileReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.tempFileReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeIoUtilWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tempFileMutex.RLock()
	defer fake.tempFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIoUtilWrapper) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.IoUtilWrapper = new(FakeIoUtilWrapper)
