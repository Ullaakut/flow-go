// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import crypto "github.com/dapperlabs/flow-go/crypto"
import flow "github.com/dapperlabs/flow-go/model/flow"
import mock "github.com/stretchr/testify/mock"

// Blocks is an autogenerated mock type for the Blocks type
type Blocks struct {
	mock.Mock
}

// ByHash provides a mock function with given fields: hash
func (_m *Blocks) ByHash(hash crypto.Hash) (*flow.Block, error) {
	ret := _m.Called(hash)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(crypto.Hash) *flow.Block); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(crypto.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ByNumber provides a mock function with given fields: number
func (_m *Blocks) ByNumber(number uint64) (*flow.Block, error) {
	ret := _m.Called(number)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(uint64) *flow.Block); ok {
		r0 = rf(number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *Blocks) Save(_a0 *flow.Block) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.Block) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
