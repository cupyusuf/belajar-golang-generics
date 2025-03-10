package belajar_golang_generics

import (
	"fmt"
	"testing"
)

func MultipleTypeParameter[T any, U any](param1 T, param2 U) {
	fmt.Println(param1, param2)
}

func TestMultipleTypeParameter(t *testing.T) {
	MultipleTypeParameter[string, int]("Yusuf", 10)
	MultipleTypeParameter[int, string](10, "Yusuf")
}
