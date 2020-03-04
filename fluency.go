package main

import (
	"reflect"
	"strconv"
	"strings"
)

// Str is the fluent object
type Str struct {
	value string
}

// String returns the string version of the fluent object
func (s *Str) String() string {
	return s.value
}

// After returns the slice of string after the first occurance of the provided value
func (s *Str) After(value string) *Str {
	index := strings.Index(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[index+len(value):]
	return s
}

// AfterLast returns the slice of the string after the last occurance of the provided value
func (s *Str) AfterLast(value string) *Str {
	index := strings.LastIndex(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[index+len(value):]
	return s
}

// Before returns the slice of the string before the first occurance of the provided value
func (s *Str) Before(value string) *Str {
	index := strings.Index(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[:index]
	return s
}

// BeforeLast returns the slice of the string before the last occurance of the provided value
func (s *Str) BeforeLast(value string) *Str {
	index := strings.LastIndex(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[:index]
	return s
}

// Append appends the give value at the end of the current fluent object
func (s *Str) Append(value interface{}) *Str {
	var appendable string

	typeof := reflect.TypeOf(value)
	switch typeof.Name() {
	case "int64":
		appendable = strconv.FormatInt(value.(int64), 10)
		break
	case "int":
		appendable = strconv.Itoa(value.(int))
	case "string":
		appendable = value.(string)
	}

	s.value += appendable
	return s
}

// On initialized the fluent object on which multiple methods can be chained
func On(initial string) *Str {
	return &Str{value: initial}
}
