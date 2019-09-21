package adapters

import (
	"context"
	"errors"
	"net/http"

	pb "simplecomputation-service/proto"
	"simplecomputation-service/service/domain"
	"simplecomputation-service/service/usecases"
)

const (
	badInputErrMsg               = "Bad input parameter specified"
	emptyInputsSpecifiedErrMsg   = "Empty inputs specified wrror message"
	failedToRunComputationErrMsg = "Failed to run computation over the inputs"
	failedToFetchResults         = "Failed to fetch results"
)

// Processor struct contains compuation interactor.
type Processor struct {
	interactor *usecases.Interactor
}

// NewProcessor takes an interactor and returns a processor.
func NewProcessor(interactor *usecases.Interactor) *Processor {
	if interactor == nil {
		return nil
	}
	return &Processor{interactor: interactor}
}

// Add adds set of inputs and return a errorResponse and an error
func (p *Processor) Add(ctx context.Context, input *pb.Input) (*pb.ErrorResponse, error) {
	if input == nil {
		return &pb.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: badInputErrMsg,
		}, errors.New(emptyInputsSpecifiedErrMsg)
	}
	domainInput := convertToDomainObject(*input)
	if err := p.interactor.Compute(domainInput); err != nil {
		return &pb.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: failedToRunComputationErrMsg,
		}, err
	}

	return &pb.ErrorResponse{Status: http.StatusOK}, nil
}

// Get results retrives the set of results after computation.
func (p *Processor) GetResults(ctx context.Context, option *pb.Option) (*pb.Result, error) {
	output, err := p.interactor.Fetch()
	if err != nil {
		return nil, errors.New(failedToFetchResults)
	}
	return &pb.Result{
		Total: int64(output.Value),
	}, nil
}

func convertToDomainObject(input pb.Input) domain.Input {
	var domainInput domain.Input
	for _, value := range input.Num {
		domainInput.Values = append(domainInput.Values, int(value))
	}
	return domainInput
}
