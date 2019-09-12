package easyconfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtraDefaultValues_BaseType(t *testing.T) {
	type Person struct {
		Name string `config:"default:joy"`
		Age  int    `config:"default:100"`
	}
	p := &Person{}
	values := ExtraDefaultValues(p)
	fmt.Println(values)

	assert.Equal(t, values["name"], "joy")
	assert.Equal(t, values["age"], 100)
}

func TestExtraDefaultValues_EmbType(t *testing.T) {
	type Son struct {
		SonName string `config:"default:ll"`
		Age     int    `config:"default:20"`
	}
	type Person struct {
		Name string `config:"default:joy"`
		Age  int    `config:"default:100"`
		Son  Son
	}
	p := &Person{}
	values := ExtraDefaultValues(p)
	fmt.Println(values)

	assert.Equal(t, values["name"], "joy")
	assert.Equal(t, values["son.son_name"], "ll")
	assert.Equal(t, values["age"], 100)
	assert.Equal(t, values["son.age"], 20)
}

func TestExtraDefaultValues_PointType(t *testing.T) {
	type Person struct {
		Name string `config:"default:joy"`
		Age  *int   `config:"default:100"`
	}
	p := &Person{}
	values := ExtraDefaultValues(p)
	fmt.Println(values)

	assert.Equal(t, values["name"], "joy")
	assert.Equal(t, values["age"], 100)
}

func TestExtraDefaultValues_PtrEmbed(t *testing.T) {
	type Son struct {
		Name string `config:"default:ll"`
		Age  int    `config:"default:20"`
	}
	type Person struct {
		Name string `config:"default:joy"`
		Age  *int   `config:"default:100"`
		Son  *Son
	}
	p := &Person{}
	values := ExtraDefaultValues(p)
	fmt.Println(values)

	assert.Equal(t, values["name"], "joy")
	assert.Equal(t, values["age"], 100)

	assert.Equal(t, values["son.name"], "ll")
	assert.Equal(t, values["son.age"], 20)
}
