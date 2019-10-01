// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"

// SessionStore is an autogenerated mock type for the SessionStore type
type SessionStore struct {
	mock.Mock
}

// AnalyticsSessionCount provides a mock function with given fields:
func (_m *SessionStore) AnalyticsSessionCount() (int64, *model.AppError) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Cleanup provides a mock function with given fields: expiryTime, batchSize
func (_m *SessionStore) Cleanup(expiryTime int64, batchSize int64) {
	_m.Called(expiryTime, batchSize)
}

// Get provides a mock function with given fields: sessionIdOrToken
func (_m *SessionStore) Get(sessionIdOrToken string) (*model.Session, *model.AppError) {
	ret := _m.Called(sessionIdOrToken)

	var r0 *model.Session
	if rf, ok := ret.Get(0).(func(string) *model.Session); ok {
		r0 = rf(sessionIdOrToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(sessionIdOrToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetSessions provides a mock function with given fields: userId
func (_m *SessionStore) GetSessions(userId string) ([]*model.Session, *model.AppError) {
	ret := _m.Called(userId)

	var r0 []*model.Session
	if rf, ok := ret.Get(0).(func(string) []*model.Session); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Session)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetSessionsWithActiveDeviceIds provides a mock function with given fields: userId
func (_m *SessionStore) GetSessionsWithActiveDeviceIds(userId string) ([]*model.Session, *model.AppError) {
	ret := _m.Called(userId)

	var r0 []*model.Session
	if rf, ok := ret.Get(0).(func(string) []*model.Session); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Session)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// PermanentDeleteSessionsByUser provides a mock function with given fields: teamId
func (_m *SessionStore) PermanentDeleteSessionsByUser(teamId string) *model.AppError {
	ret := _m.Called(teamId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: sessionIdOrToken
func (_m *SessionStore) Remove(sessionIdOrToken string) *model.AppError {
	ret := _m.Called(sessionIdOrToken)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(sessionIdOrToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// RemoveAllSessions provides a mock function with given fields:
func (_m *SessionStore) RemoveAllSessions() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Save provides a mock function with given fields: session
func (_m *SessionStore) Save(session *model.Session) (*model.Session, *model.AppError) {
	ret := _m.Called(session)

	var r0 *model.Session
	if rf, ok := ret.Get(0).(func(*model.Session) *model.Session); ok {
		r0 = rf(session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.Session) *model.AppError); ok {
		r1 = rf(session)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateDeviceId provides a mock function with given fields: id, deviceId, expiresAt
func (_m *SessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) (string, *model.AppError) {
	ret := _m.Called(id, deviceId, expiresAt)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, int64) string); ok {
		r0 = rf(id, deviceId, expiresAt)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, int64) *model.AppError); ok {
		r1 = rf(id, deviceId, expiresAt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateLastActivityAt provides a mock function with given fields: sessionId, time
func (_m *SessionStore) UpdateLastActivityAt(sessionId string, time int64) *model.AppError {
	ret := _m.Called(sessionId, time)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int64) *model.AppError); ok {
		r0 = rf(sessionId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateRoles provides a mock function with given fields: userId, roles
func (_m *SessionStore) UpdateRoles(userId string, roles string) (string, *model.AppError) {
	ret := _m.Called(userId, roles)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(userId, roles)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(userId, roles)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}