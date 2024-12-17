package calculator

import (
	"errors"
	"math"
	"strconv"

	"github.com/wDRxxx/yandex_calculator_api/internal/service"
	"github.com/wDRxxx/yandex_calculator_api/pkg/stack"
)

type calculator struct {
	priority map[byte]int
}

func NewCalculatorService() service.CalculatorService {
	var m = map[byte]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
	}

	return &calculator{
		priority: m,
	}
}

func (c *calculator) execOp(op byte, a float64, b float64) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	case '^':
		return math.Pow(a, b)
	}

	return 0
}

func (c *calculator) ConvertToPolishNotation(expression string) (string, error) {
	result := ""
	st := stack.NewStack[byte]()

	for i := 0; i < len(expression); i++ {
		if expression[i] >= '0' && expression[i] <= '9' {
			num := ""

			for i < len(expression) && ((expression[i] >= '0' && expression[i] <= '9') || expression[i] == '.') {
				num += string(expression[i])
				i++
			}

			result += num + "|"
			if i != len(expression)-1 {
				i--
			}
		} else if expression[i] == '(' {
			st.Push(expression[i])
		} else if expression[i] == ')' {
			for st.Len() > 0 && st.Peek() != '(' {
				result += string(st.Pop())
			}
			st.Pop()
		} else if _, ok := c.priority[expression[i]]; ok {
			for st.Len() > 0 && (c.priority[st.Peek()] >= c.priority[expression[i]]) {
				result += string(st.Pop())
			}
			st.Push(expression[i])
		} else {
			return "", errors.New("wrong expression format")
		}
	}

	for st.Len() > 0 {
		result += string(st.Pop())
	}

	return result, nil
}

func (c *calculator) Calculate(input string) (float64, error) {
	expression, err := c.ConvertToPolishNotation(input)
	if err != nil {
		return 0, err
	}

	nums := stack.NewStack[float64]()

	for i := 0; i < len(expression); i++ {
		if expression[i] >= '0' && expression[i] <= '9' {
			num := ""

			for i < len(expression) && ((expression[i] >= '0' && expression[i] <= '9') || expression[i] == '.') {
				num += string(expression[i])
				i++
			}

			n, err := strconv.ParseFloat(num, 64)
			if err != nil {
				return 0, err
			}

			nums.Push(n)
		} else if _, ok := c.priority[expression[i]]; ok {
			if nums.Len() > 1 {
				b := nums.Pop()
				a := nums.Pop()

				nums.Push(c.execOp(expression[i], a, b))
			}
		}
	}

	return nums.Pop(), nil
}
