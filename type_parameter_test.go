package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Hello World")
	assert.Equal(t, "Hello World", result)

	var resultNumber int = Length[int](10)
	assert.Equal(t, 10, resultNumber)
}
