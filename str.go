package main

import "strings"

type str struct {
	value string
}

func (s *str) toString() string {
	return s.value
}

func (s *str) after(value string) *str {
	index := strings.Index(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[index+len(value):]
	return s
}

func (s *str) afterLast(value string) *str {
	index := strings.LastIndex(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[index+len(value):]
	return s
}

func (s *str) before(value string) *str {
	index := strings.Index(s.value, value)
	if index == -1 {
		return s
	}

	s.value = s.value[:index]
	return s
}

func (s *str) append(value string) *str {
	s.value += value
	return s
}

func of(initial string) *str {
	return &str{value: initial}
}
