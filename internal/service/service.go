package service

type CalculatorService interface {
	ConvertToPolishNotation(expression string) (string, error)
	Calculate(expression string) (float64, error)
}
