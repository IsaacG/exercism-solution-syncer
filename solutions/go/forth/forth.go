package forth

import (
	"fmt"
	"strconv"
	"strings"
)

// emulator contains emulator details: the stack and any custom definitions.
type emulator struct {
	stack     []int
	operators map[string][]string
}

// define handles a ": name ... ;" definition.
func (e *emulator) define(line []string) error {
	length := len(line)
	if length < 4 || line[0] != ":" || line[length-1] != ";" {
		return fmt.Errorf("malformed define, %v", line)
	}
	// The custom command name cannot be an integer.
	if _, err := strconv.Atoi(line[1]); err == nil {
		return fmt.Errorf("invalid define %s", line[1])
	}
	var operation []string
	for _, word := range line[2 : length-1] {
		// When storing a definition, expand existing definitions.
		if expanded, ok := e.operators[word]; ok {
			operation = append(operation, expanded...)
		} else {
			operation = append(operation, word)
		}
	}
	e.operators[line[1]] = operation
	return nil
}

// Evalute a command that operates on the stack.
func (e *emulator) stackOp(op string, count int) error {
	length := len(e.stack)
	if length < count {
		return fmt.Errorf("%s requires %d values; stack size is %d", op, count, length)
	}
	// Read the operands from the stack.
	var result []int
	operands := make([]int, count)
	for i := range count {
		operands[i] = e.stack[length-count+i]
	}
	switch op {
	case "+":
		result = []int{operands[0] + operands[1]}
	case "-":
		result = []int{operands[0] - operands[1]}
	case "*":
		result = []int{operands[0] * operands[1]}
	case "/":
		if operands[1] == 0 {
			return fmt.Errorf("illegal division by zero")
		}
		result = []int{operands[0] / operands[1]}
	case "DUP":
		result = []int{operands[0], operands[0]}
	case "DROP":
		result = []int{}
	case "SWAP":
		result = []int{operands[1], operands[0]}
	case "OVER":
		result = []int{operands[0], operands[1], operands[0]}
	default:
		panic("Invalid op, " + op)
	}
	// Update the stack.
	e.stack = append(e.stack[:length-count], result...)
	return nil
}

// eval evaluates one line.
func (e *emulator) eval(line []string) error {
	if line[0] == ":" {
		return e.define(line)
	}
	for _, word := range line {
		var err error
		if val, convErr := strconv.Atoi(word); convErr == nil {
			// Integers get pushed to the stack.
			e.stack = append(e.stack, val)
		} else if operation, ok := e.operators[word]; ok {
			// Custom definitions take precendence over built-in operators.
			err = e.eval(operation)
		} else {
			switch word {
			case "+", "-", "*", "/":
				err = e.stackOp(word, 2)
			case "DUP", "DROP":
				err = e.stackOp(word, 1)
			case "SWAP", "OVER":
				err = e.stackOp(word, 2)
			default:
				return fmt.Errorf("invalid command %s", word)
			}
		}
		// Stop on an error.
		if err != nil {
			return err
		}
	}
	return nil
}

// Forth emulates a Forth program.
func Forth(input []string) ([]int, error) {
	em := &emulator{nil, map[string][]string{}}
	for _, line := range input {
		if err := em.eval(strings.Fields(strings.ToUpper(line))); err != nil {
			return nil, err
		}
	}
	return em.stack, nil
}
