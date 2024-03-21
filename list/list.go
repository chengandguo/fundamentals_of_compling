package list

import (
	"errors"
)

type StringList []string

// get peek element of string list
func (l *StringList) Peek(args ...string) (string, error) {
	length := len(*l)
	if length > 0 {
		target := (*l)[length-1]
		return target, nil
	} else {
		return "", errors.New("list is empty")
	}
}

func (l *StringList) Push(args ...string) int {
	*l = append(*l, args...)
	return len(*l)
}

func (l *StringList) Pop() (string, error) {
	length := len(*l)
	if length > 0 {
		target := (*l)[length-1]
		*l = (*l)[:length-1]
		return target, nil
	} else {
		return "", errors.New("list is empty")
	}
}

func (l *StringList) Shift() (string, error) {
	length := len(*l)
	if length > 0 {
		target := (*l)[0]
		*l = (*l)[1:length]
		return target, nil
	} else {
		return "", errors.New("list is empty")
	}
}

func (l *StringList) UnShift(args ...string) int {
	newL := StringList{}
	for _, item := range args {
		newL = append(newL, item)
	}
	*l = append(newL, (*l)...)
	return len(*l)
}

func (l *StringList) Find(findFunc func(value string) bool) (string, error) {
	for _, item := range *l {
		if findFunc(item) {
			return item, nil
		}
	}
	return "", errors.New("can't find target value")
}

func (l *StringList) FindIndex(findFunc func(value string) bool) (int, error) {
	for index, item := range *l {
		if findFunc(item) {
			return index, nil
		}
	}
	return -1, errors.New("can't find target value")
}

func (l *StringList) Includes(value string) bool {
	for _, item := range *l {
		if item == value {
			return true
		}
	}
	return false
}

func (l *StringList) Map(mapFunc func(value string, index int, thisArg *StringList) any) []any {
	list := []any{}
	for index, item := range *l {
		list = append(list, mapFunc(item, index, l))
	}
	return list
}
