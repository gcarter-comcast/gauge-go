package testsuit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldRunImplementation(t *testing.T) {
	called := false
	var calledWith interface{}
	step := Step{
		Description: "Test description",
		Impl: func(args ...interface{}) {
			calledWith = args
			called = true
		},
	}

	step.Execute(1, true, "foo")

	assert.True(t, called)
	assert.Contains(t, calledWith, 1)
	assert.Contains(t, calledWith, true)
	assert.Contains(t, calledWith, "foo")
}

func TestShouldReturnPassedMethodExecutionResult(t *testing.T) {
	var called bool
	step := Step{
		Description: "Test description",
		Impl: func(args ...interface{}) {
			called = true
		},
	}

	res := step.Execute("foo")

	assert.True(t, called)
	assert.False(t, res.GetFailed())
	assert.NotZero(t, res.GetExecutionTime())
	assert.Zero(t, res.GetErrorMessage())
	assert.Zero(t, res.GetStackTrace())
}

func TestShouldReturnFailedMethodExecutionResult(t *testing.T) {
	var called bool
	step := Step{
		Description: "Test description",
		Impl: func(args ...interface{}) {
			called = true
			var a []string
			fmt.Println(a[7])
		},
	}

	res := step.Execute("foo")

	assert.True(t, called)
	assert.True(t, res.GetFailed())
	assert.NotZero(t, res.GetExecutionTime())
	assert.Equal(t, "runtime error: index out of range", res.GetErrorMessage())
	assert.NotZero(t, res.GetStackTrace())
}
