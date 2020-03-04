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

func TestCanContains(t *testing.T) {
	yes := On("Water in a glass").Contains("Water")
	no := On("Water in a glass").Contains("Juice")
	assert.True(t, yes)
	assert.False(t, no)
}

func TestCanContainsAny(t *testing.T) {
	yes := On("Water in a glass").ContainsAny([]string{"Juice", "Water"})
	assert.True(t, yes)

	no := On("Water in a glass").ContainsAny([]string{"Soda", "Juice"})
	assert.False(t, no)
}

func TestCanContainsAll(t *testing.T) {
	yes := On("Water in a glass").ContainsAll([]string{"glass", "Water"})
	assert.True(t, yes)

	no := On("Water in a glass").ContainsAll([]string{"Water", "Juice"})
	assert.False(t, no)
}

func TestCanEndsWith(t *testing.T) {
	yes := On("A complete sentence.").EndsWith(".")
	assert.True(t, yes)

	no := On("Is this a question?").EndsWith(".")
	assert.False(t, no)
}

func TestCanEquals(t *testing.T) {
	yes := On("Something").Equals("Something")
	assert.True(t, yes)

	no := On("Nothing").Equals("Something")
	assert.False(t, no)
}

func TestCanFinish(t *testing.T) {
	s := On("some/path/to/somewhere").Finish("/")
	assert.Equal(t, "some/path/to/somewhere/", s.String())

	s2 := On("some/path/to/somewhere/").Finish("/")
	assert.Equal(t, "some/path/to/somewhere/", s2.String())
}

func TestCanSplit(t *testing.T) {
	splits := On("some/path/to/somewhere").Split("/")
	assert.Len(t, splits, 4)
	assert.Contains(t, splits, "some")
	assert.Contains(t, splits, "path")
	assert.Contains(t, splits, "to")
	assert.Contains(t, splits, "somewhere")
}

func TestCanStartsWith(t *testing.T) {
	yes := On("A complete sentence.").StartsWith("A comple")
	assert.True(t, yes)

	no := On("Is this a question?").StartsWith("A comple")
	assert.False(t, no)
}

func TestCanTrim(t *testing.T) {
	s := On("   space around a rocket   ").Trim()
	assert.Equal(t, "space around a rocket", s.String())
}
