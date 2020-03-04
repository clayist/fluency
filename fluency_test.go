package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFluencyIsInitialized(t *testing.T) {
	s := On("")
	assert.NotNil(t, s)
	assert.IsType(t, &Str{}, s)
}

func TestFluentObjectCanBeCastedIntoString(t *testing.T) {
	input := "Hello"
	s := On(input)
	assert.Equal(t, fmt.Sprint(s), s.String())
	assert.Equal(t, input, fmt.Sprint(s))
}

func TestCanAppendString(t *testing.T) {
	s := On("Hello").Append("World")
	assert.Equal(t, "HelloWorld", s.String())
}

func TestCanAppendMultipleTimes(t *testing.T) {
	s := On("Hello").Append(" ").Append("World")
	assert.Equal(t, "Hello World", s.String())
}

func TestCanAppendInt(t *testing.T) {
	s := On("id-").Append(123)
	assert.Equal(t, "id-123", s.String())
}

func TestCanAppendInt64(t *testing.T) {
	timestamp := time.Now().Unix()
	s := On("timestamp-").Append(time.Now().Unix())
	assert.Equal(t, fmt.Sprintf("timestamp-%s", strconv.FormatInt(timestamp, 10)), s.String())
}

func TestCanAfter(t *testing.T) {
	s := On("to be or not to be").After("to ")
	assert.Equal(t, "be or not to be", s.String())
}

func TestCanAfterLast(t *testing.T) {
	s := On("to be or not to be").AfterLast("to ")
	assert.Equal(t, "be", s.String())
}

func TestCanBefore(t *testing.T) {
	s := On("to be or not to be").Before("to ")
	assert.Equal(t, "", s.String())
}

func TestCanBeforeLast(t *testing.T) {
	s := On("to be or not to be").BeforeLast(" to")
	assert.Equal(t, "to be or not", s.String())
}
