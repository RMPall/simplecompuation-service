package adapters

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simplecomputation-service/proto"
	"simplecomputation-service/service/adapters/mocks"
	"simplecomputation-service/service/usecases"
)

func TestNewProcessor_WhenInteractorIsNil_ReturnsNilProcessor(t *testing.T) {
	t.Parallel()

	assert.Nil(t, NewProcessor(nil))
}

func TestNewProcessor_ReturnsProcessor(t *testing.T) {
	t.Parallel()
	interactor := &usecases.Interactor{}

	assert.NotNil(t, NewProcessor(interactor))
}

func TestAdd_WhenInputIsNil_ReturnsError(t *testing.T) {
	t.Parallel()
	interactor := &usecases.Interactor{}
	processor := NewProcessor(interactor)

	resp, err := processor.Add(context.TODO(), nil)

	assert.NotNil(t, err)
	assert.Equal(t, int64(http.StatusBadRequest), resp.Status)
}

func TestAdd_WhenComputationFails_ReturnsError(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := usecases.NewInteractor(repoMock)
	processor := NewProcessor(interactor)
	expectedErrMsg := "something bad happened"
	repoMock.On("Save", mock.Anything).Return(errors.New(expectedErrMsg))

	resp, err := processor.Add(context.TODO(), &compute.Input{})

	assert.NotNil(t, err)
	assert.Equal(t, int64(http.StatusBadRequest), resp.Status)
}

func TestAdd_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := usecases.NewInteractor(repoMock)
	processor := NewProcessor(interactor)
	repoMock.On("Save", mock.Anything).Return(nil)

	resp, err := processor.Add(context.TODO(), &compute.Input{})

	assert.Nil(t, err)
	assert.Equal(t, int64(http.StatusOK), resp.Status)
}

func TestGetResults_WhenFailedToFetchResults_ReturnError(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := usecases.NewInteractor(repoMock)
	processor := NewProcessor(interactor)
	expectedErrMsg := "something bad happened"
	repoMock.On("Get").Return(int(0), errors.New(expectedErrMsg))

	_, err := processor.GetResults(context.TODO(), &compute.Option{})

	assert.NotNil(t, err)
}

func TestGetResults_ReturnsSuccess(t *testing.T) {
	t.Parallel()
	repoMock := &mocks.RepoMock{}
	interactor := usecases.NewInteractor(repoMock)
	processor := NewProcessor(interactor)
	repoMock.On("Get").Return(int(10), nil)

	_, err := processor.GetResults(context.TODO(), &compute.Option{})

	assert.Nil(t, err)
}
