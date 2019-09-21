package mocks

import "github.com/stretchr/testify/mock"

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) Save(total int) error {
	args := r.Called(total)
	return args.Error(0)
}

func (r *RepoMock) Get() (int, error) {
	args := r.Called()
	return args.Get(0).(int), args.Error(1)
}
