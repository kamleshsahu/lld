package command

type ICommand interface {
	IsValid(token string) bool
	Execute(token string, low, high int) []int
}
