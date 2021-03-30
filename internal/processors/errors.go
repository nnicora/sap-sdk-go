package processors

type ErrorChecker interface {
	HasError() bool
}
