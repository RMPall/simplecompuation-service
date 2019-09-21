package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simplecompuation-service/service/domain"
	"simplecompuation-service/service/usecases/mocks"
)

func TestNewInteractor_WhenRepoIsNil_ReturnsNilInteractor(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NewInteractor(nil))
}

func TestNewInteractor_ReturnsInteractor(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}

	assert.NotNil(t, NewInteractor(repoMock))
}

func TestCompute_WhenResultCannotbeSaved_ReturnsError(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := NewInteractor(repoMock)
	expectedErrMsg := "something bad happened"
	repoMock.On("Save", mock.Anything).Return(errors.New(expectedErrMsg))
	err := interactor.Compute(domain.Input{})

	assert.NotNil(t, err)
}

func TestCompute_ReturnsNilError(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := NewInteractor(repoMock)
	repoMock.On("Save", mock.Anything).Return(nil)
	err := interactor.Compute(domain.Input{})

	assert.Nil(t, err)
}

func TestFetch_WhenFailedToGetResults_ReturnsError(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	expectedErrMsg := "something bad happened"
	interactor := NewInteractor(repoMock)
	repoMock.On("Get").Return(0, errors.New(expectedErrMsg))

	_, err := interactor.Fetch()

	assert.NotNil(t, err)
}

func TestFetch_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := NewInteractor(repoMock)
	repoMock.On("Get").Return(10, nil)

	_, err := interactor.Fetch()

	assert.Nil(t, err)
}
