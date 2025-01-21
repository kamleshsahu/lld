package operator

type IOperator interface {
	IsApplicable(token string) bool
	Execute(token string, low, high int) ([]int, error)
}

func DefaultOperatorList() []IOperator {
	commands := []IOperator{
		NewWildCardParser(),
		NewNumberParser(),
		NewStepParser(),
		NewRangeParser(),
		NewCommaParser(),
	}
	return commands
}
