package usecases

// Repository consists of Save the result of a computsation and fetching the results of a compuation.
type Repository interface {
	Save(total int) error
	Get() (int, error)
}
