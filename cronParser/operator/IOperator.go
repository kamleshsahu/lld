package operator

type IOperator interface {
	IsApplicable(token string) bool
	Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error)
}

func DefaultOperatorList() []IOperator {
	commands := []IOperator{
		NewWildCardParser(),
		NewStepParser(),
		NewCommaParser(),
	}
	return commands
}
