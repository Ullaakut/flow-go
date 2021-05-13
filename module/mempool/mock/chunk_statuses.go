// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	verification "github.com/onflow/flow-go/model/verification"
)

// ChunkStatuses is an autogenerated mock type for the ChunkStatuses type
type ChunkStatuses struct {
	mock.Mock
}

// Add provides a mock function with given fields: status
func (_m *ChunkStatuses) Add(status *verification.ChunkStatus) bool {
	ret := _m.Called(status)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*verification.ChunkStatus) bool); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *ChunkStatuses) All() []*verification.ChunkStatus {
	ret := _m.Called()

	var r0 []*verification.ChunkStatus
	if rf, ok := ret.Get(0).(func() []*verification.ChunkStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*verification.ChunkStatus)
		}
	}

	return r0
}

// ByID provides a mock function with given fields: chunkID
func (_m *ChunkStatuses) ByID(chunkID flow.Identifier) (*verification.ChunkStatus, bool) {
	ret := _m.Called(chunkID)

	var r0 *verification.ChunkStatus
	if rf, ok := ret.Get(0).(func(flow.Identifier) *verification.ChunkStatus); ok {
		r0 = rf(chunkID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*verification.ChunkStatus)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(chunkID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Rem provides a mock function with given fields: chunkID
func (_m *ChunkStatuses) Rem(chunkID flow.Identifier) bool {
	ret := _m.Called(chunkID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(chunkID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
