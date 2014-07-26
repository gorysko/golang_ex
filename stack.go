package stack

import "errors"

type  Stack []interface{}

func (stack Stack) Len() int {
    return len(stack)
}

func (stack Stack) Cap() int {
    return cap(stack)
}

func (stack *Stack) Push(x interface{}) {
    *stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
    if len(stack) == 0 {
        return nil, errors.New("can't get last in empty stack")
    }

    return stack[len(stack) - 1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
    temp_stack := *stack
    if len(temp_stack) == 0 {
        return nil, errors.New("empty stacj can't be poped")
    }

    x := temp_stack[len(temp_stack) - 1]
    *stack = temp_stack[:len(temp_stack) - 1]
    return x, nil
}


