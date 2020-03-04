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

// String returns the string version
// of the fluent object
func (s *Str) String() string {
	return s.value
}

// After returns the slice of string after the
// first occurance of the provided value
func (s *Str) After(value string) *Str {
	index := strings.Index(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[index+len(value):]
	return s
}

// AfterLast returns the slice of the string after
// the last occurance of the provided value
func (s *Str) AfterLast(value string) *Str {
	index := strings.LastIndex(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[index+len(value):]
	return s
}

// Before returns the slice of the string before
// the first occurance of the provided value
func (s *Str) Before(value string) *Str {
	index := strings.Index(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[:index]
	return s
}

// BeforeLast returns the slice of the string before
// the last occurance of the provided value
func (s *Str) BeforeLast(value string) *Str {
	index := strings.LastIndex(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[:index]
	return s
}

// Append appends the give value at the
// end of the current fluent object
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

// Contains returns a boolean signifying whether
// the fluent object contains the given value
func (s *Str) Contains(value string) bool {
	return strings.Contains(s.value, value)
}

// ContainsAny returns a boolean signifying whether
// the fluent object contains any of the given values
func (s *Str) ContainsAny(values []string) bool {
	for _, value := range values {
		if s.Contains(value) {
			return true
		}
	}
	return false
}

// ContainsAll returns a boolean signifying whether
// the fluent object contains all of the given values
func (s *Str) ContainsAll(values []string) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

// EndsWith returns a boolean signifying whether
// the fluent object ends with the given value
func (s *Str) EndsWith(value string) bool {
	return strings.HasSuffix(s.value, value)
}

// Equals returns a boolean signifying whether
// the fluent object is equals to the given value
func (s *Str) Equals(value string) bool {
	return s.value == value
}

// Finish ends the fluent object with the given value
func (s *Str) Finish(value string) *Str {
	if s.EndsWith(value) {
		return s
	}

	return s.Append(value)
}

// Split breaks the fluent object into a slice of strings
// at every occurance of the separator
func (s *Str) Split(separator string) []string {
	return strings.Split(s.value, separator)
}

// StartsWith returns a boolean signifying whether
// the fluent object starts with the given value
func (s *Str) StartsWith(value string) bool {
	return strings.HasPrefix(s.value, value)
}

// Trim removes the leading and trailing spaces
func (s *Str) Trim() *Str {
	s.value = strings.TrimSpace(s.value)
	return s
}

// On initialized the fluent object on which multiple methods can be chained
func On(initial string) *Str {
	return &Str{value: initial}
}
