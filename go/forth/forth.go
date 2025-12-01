package forth

import (
	"errors"
	"strconv"
	"strings"
)

// each operation modifies the stack
var operations = map[string]func(stack *[]int) bool{
	"dup": func(stack *[]int) bool {
		if len(*stack) < 1 {
			return false
		}
		top := (*stack)[len(*stack)-1]
		*stack = append(*stack, top)
		return true
	},

	"drop": func(stack *[]int) bool {
		if len(*stack) < 1 {
			return false
		}
		pop(stack)
		return true
	},

	"swap": func(stack *[]int) bool {
		if len(*stack) < 2 {
			return false
		}
		b := pop(stack)
		a := pop(stack)
		*stack = append(*stack, b, a)
		return true
	},

	"over": func(stack *[]int) bool {
		if len(*stack) < 2 {
			return false
		}
		s := *stack
		*stack = append(*stack, s[len(s)-2])
		return true
	},

	"+": func(stack *[]int) bool {
		if len(*stack) < 2 {
			return false
		}
		b := pop(stack)
		a := pop(stack)
		*stack = append(*stack, a+b)
		return true
	},

	"-": func(stack *[]int) bool {
		if len(*stack) < 2 {
			return false
		}
		b := pop(stack)
		a := pop(stack)
		*stack = append(*stack, a-b)
		return true
	},

	"*": func(stack *[]int) bool {
		if len(*stack) < 2 {
			return false
		}
		b := pop(stack)
		a := pop(stack)
		*stack = append(*stack, a*b)
		return true
	},

	"/": func(stack *[]int) bool {
		if len(*stack) < 2 {
			return false
		}
		b := pop(stack)
		if b == 0 {
			return false
		}
		a := pop(stack)
		*stack = append(*stack, a/b)
		return true
	},
}

func Forth(input []string) ([]int, error) {
	var stack []int
	words := map[string][]string{}

	for _, s := range input {
		ops := strings.Fields(strings.ToLower(s))
		if len(ops) == 0 {
			continue
		}

		if ops[0] == ":" {
			err := defineWord(ops, words)
			if err != nil {
				return stack, err
			}
			continue
		}

		if err := runOps(ops, &stack, words); err != nil {
			return stack, err
		}
	}

	return stack, nil
}

func defineWord(ops []string, words map[string][]string) error {
	if len(ops) < 3 || ops[len(ops)-1] != ";" {
		return errors.New("invalid words")
	}

	name := ops[1]

	if _, ok := toInt(name); ok {
		return errors.New("invalid words")
	}

	rawBody := ops[2 : len(ops)-1]
	var expanded []string
	for _, t := range rawBody {
		if def, ok := words[t]; ok {
			expanded = append(expanded, def...)
		} else {
			expanded = append(expanded, t)
		}
	}

	words[name] = expanded
	return nil
}

func runOps(ops []string, stack *[]int, words map[string][]string) error {
	for _, op := range ops {
		if num, ok := toInt(op); ok {
			*stack = append(*stack, num)
			continue
		}

		if body, ok := words[op]; ok {
			if err := runOps(body, stack, words); err != nil {
				return err
			}
			continue
		}

		if fn, ok := operations[op]; ok {
			if !fn(stack) {
				return errors.New("invalid operation")
			}
			continue
		}

		return errors.New("invalid operation")
	}

	return nil
}

func toInt(op string) (int, bool) {
	if len(op) == 0 {
		return 0, false
	}
	num, err := strconv.Atoi(op)
	if err != nil {
		return 0, false
	}
	return num, true
}

func pop(stack *[]int) int {
	s := *stack
	x := s[len(s)-1]
	*stack = s[:len(s)-1]
	return x
}
