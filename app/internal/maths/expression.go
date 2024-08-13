package maths

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value       string `json:"value"`
	IsOperation bool   `json:"is_operation"`
}

type Expression struct {
	Elements []Node `json:"elements"`
	opStack  OpStack
	Result   string `json:"result"`
	Error    error  `json:"error"`
}

type OpStack struct {
	elements []Node
}

func (e *Expression) Parse(input string) {
	operand := ""

	for i, c := range input {
		c := string(c)
		if e.IsOp(c) {
			op := string(c)
			e.Elements = append(e.Elements, Node{Value: operand, IsOperation: false})

			for len(e.opStack.elements) != 0 && e.Order(e.opStack.Peek().Value) >= e.Order(op) {
				e.Elements = append(e.Elements, e.opStack.Pop())
			}
			e.opStack.elements = append(e.opStack.elements, Node{Value: op, IsOperation: true})
			operand = ""
		} else {
			operand += string(c)
			if i == len(input)-1 {
				e.Elements = append(e.Elements, Node{Value: operand, IsOperation: false})
			}
		}
	}

	for len(e.opStack.elements) != 0 {
		e.Elements = append(e.Elements, e.opStack.Pop())
	}
}

func (e *Expression) Solve() {
	tmpStack := OpStack{}
	for _, n := range e.Elements {
		if !n.IsOperation {
			tmpStack.elements = append(tmpStack.elements, n)
			continue
		}

		one := tmpStack.Pop()
		two := tmpStack.Pop()

		tmpStack.elements = append(tmpStack.elements, e.calc(n, one, two))
	}

	e.Result = tmpStack.Pop().Value
}

func (e *Expression) calc(op Node, one Node, two Node) Node {
	val1, err1 := strconv.ParseFloat(one.Value, 64)
	val2, err2 := strconv.ParseFloat(two.Value, 64)

	if err1 != nil {
		e.Error = err1
	}

	if err2 != nil {
		e.Error = err1
	}

	var result float64

	switch op.Value {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 == 0.0 {
			e.Error = fmt.Errorf("cannot divide by 0")
			return Node{}
		}
		result = val1 / val2
	}

	resStr := ""

	if result == float64(int64(result)) {
		resStr = strconv.FormatInt(int64(result), 10)
	} else {
		resStr = strconv.FormatFloat(result, 'f', -1, 64)
	}

	return Node{Value: resStr}
}

func (e *Expression) IsOp(c string) bool {
	return (c == "+" || c == "-" || c == "*" || c == "/")
}

func (e *Expression) Order(value string) int {
	if value == "+" || value == "-" {
		return 1
	}
	return 2
}

func (o *OpStack) Peek() Node {
	if len(o.elements) == 0 {
		return Node{}
	}
	return o.elements[len(o.elements)-1]
}

func (o *OpStack) Pop() Node {
	node := o.Peek()
	lastPos := len(o.elements) - 1
	if lastPos < 0 {
		return node
	}
	o.elements = append(o.elements[:lastPos], o.elements[lastPos+1:]...)
	return node
}

func (o *OpStack) Remove(slice []Node, s int) []Node {
	return append(slice[:s], slice[s+1:]...)
}
