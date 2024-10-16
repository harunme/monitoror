// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/monitorables/github/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetChecks provides a mock function with given fields: owner, repository, ref
func (_m *Repository) GetChecks(owner string, repository string, ref string) (*models.Checks, error) {
	ret := _m.Called(owner, repository, ref)

	if len(ret) == 0 {
		panic("no return value specified for GetChecks")
	}

	var r0 *models.Checks
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*models.Checks, error)); ok {
		return rf(owner, repository, ref)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Checks); ok {
		r0 = rf(owner, repository, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Checks)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repository, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommit provides a mock function with given fields: owner, repository, sha
func (_m *Repository) GetCommit(owner string, repository string, sha string) (*models.Commit, error) {
	ret := _m.Called(owner, repository, sha)

	if len(ret) == 0 {
		panic("no return value specified for GetCommit")
	}

	var r0 *models.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*models.Commit, error)); ok {
		return rf(owner, repository, sha)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Commit); ok {
		r0 = rf(owner, repository, sha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repository, sha)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields: query
func (_m *Repository) GetCount(query string) (int, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for GetCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPullRequest provides a mock function with given fields: owner, repository, id
func (_m *Repository) GetPullRequest(owner string, repository string, id int) (*models.PullRequest, error) {
	ret := _m.Called(owner, repository, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPullRequest")
	}

	var r0 *models.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int) (*models.PullRequest, error)); ok {
		return rf(owner, repository, id)
	}
	if rf, ok := ret.Get(0).(func(string, string, int) *models.PullRequest); ok {
		r0 = rf(owner, repository, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(owner, repository, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPullRequests provides a mock function with given fields: owner, repository
func (_m *Repository) GetPullRequests(owner string, repository string) ([]models.PullRequest, error) {
	ret := _m.Called(owner, repository)

	if len(ret) == 0 {
		panic("no return value specified for GetPullRequests")
	}

	var r0 []models.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]models.PullRequest, error)); ok {
		return rf(owner, repository)
	}
	if rf, ok := ret.Get(0).(func(string, string) []models.PullRequest); ok {
		r0 = rf(owner, repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
