package operator

type IOperator interface {
	IsApplicable(token string) bool
	Execute(token string, low, high int, toNumber func(val string) (int, error)) ([]int, error)
}

func DefaultOperatorList() []IOperator {
	commands := []IOperator{
		NewWildCardParser(),
		NewSingleValueParser(),
		NewStepParser(),
		NewRangeParser(),
		NewCommaParser(),
	}
	return commands
}
