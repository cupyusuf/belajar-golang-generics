package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Yusuf",
		Second: "Supriadi",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Yusuf",
		Second: "Supriadi",
	}

	assert.Equal(t, "Firlana", data.ChangeFirst("Firlana"))
	assert.Equal(t, "Hello Yusuf", data.SayHello("Yusuf"))
}
