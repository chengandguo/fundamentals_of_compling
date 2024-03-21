package main

import (
	"fmt"
	"monkey/list"
)

func main() {
	fmt.Println("hello world")
	str3 := []string{"(", "1", "+", "2", ")", "*", "3", "-", "4", "*", "5"}
	fmt.Println(infixToPostfixNotation(str3))
}

func infixToPostfixNotation(str []string) []string {
	priorityMap := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"(": 3,
		")": 3,
	}
	var operators list.StringList
	var nums list.StringList

	for _, item := range str {
		currentPriority, isOperator := priorityMap[item]
		peekOperator, error := operators.Peek()
		if isOperator {
			if len(operators) == 0 || item == "(" || peekOperator == "(" {
				operators.Push(item)
			} else if item == ")" {
				for op, err := operators.Pop(); op != "(" && err == nil; {
					nums.Push(op)
					op, err = operators.Pop()
				}
			} else {
				if error != nil {
					panic(error)
				}
				peekPriority := priorityMap[peekOperator]
				if peekPriority < currentPriority {
					operators.Push(item)
				} else {
					for peekPriority >= currentPriority {
						op, error := operators.Pop()
						if error != nil {
							break
						}
						nums.Push(op)
						peekOperator, error := operators.Peek()
						if error != nil {
							break
						}
						peekPriority = priorityMap[peekOperator]
					}
					operators.Push(item)
				}
			}
		} else {
			nums.Push(item)
		}
	}

	for op, err := operators.Pop(); err == nil; op, err = operators.Pop() {
		nums.Push(op)
	}

	return nums
}
