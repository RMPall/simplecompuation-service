package usecases

import (
	"simplecomputation-service/service/domain"
)

// Interactor consists of repository to save the compuation
type Interactor struct {
	repo Repository
}

// NewInteractor takes repo and return interactor
func NewInteractor(repo Repository) *Interactor {
	if repo == nil {
		return nil
	}
	return &Interactor{repo: repo}
}

// Compute takes input and returns error
func (i *Interactor) Compute(input domain.Input) error {
	var total int
	for _, value := range input.Values {
		total = total + value
	}

	if err := i.repo.Save(total); err != nil {
		return err
	}

	return nil
}

// Fetch fetches the results
func (i *Interactor) Fetch() (*domain.Output, error) {
	output, err := i.repo.Get()
	if err != nil {
		return nil, err
	}

	return &domain.Output{Value: output}, nil
}
