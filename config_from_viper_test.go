package easyconfig

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigFromViper_BaseType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := &Person{}

	v := viper.New()
	var defaultValues map[string]interface{} = map[string]interface{}{
		"name": "joy",
		"age":  100,
	}
	SetViperDefault(v, defaultValues)
	err := LoadConfigFromViper(p, v)
	fmt.Println(p)
	assert.NoError(t, err)
	assert.Equal(t, p.Name, "joy")
	assert.Equal(t, p.Age, 100)
}

func TestLoadConfigFromViper_Embed(t *testing.T) {
	type Son struct {
		Name string `config:"default:ll"`
		Age  int    `config:"default:20"`
	}
	type Person struct {
		Name string `config:"default:joy"`
		Age  int    `config:"default:100"`
		Son  Son
	}
	p := &Person{}

	v := viper.New()
	var defaultValues map[string]interface{} = map[string]interface{}{
		"name":     "joy",
		"age":      100,
		"son.name": "ss",
		"son.age":  20,
	}
	SetViperDefault(v, defaultValues)

	err := LoadConfigFromViper(p, v)

	assert.NoError(t, err)
	fmt.Printf("%+v \n", p)
}

func TestLoadConfigFromViper_Pointer(t *testing.T) {
	type Person struct {
		Name *string `config:"default:joy"`
		Age  *int    `config:"default:100"`
	}
	p := &Person{}

	v := viper.New()
	var defaultValues = map[string]interface{}{
		"name": "joy",
		"age":  100,
	}
	SetViperDefault(v, defaultValues)

	err := LoadConfigFromViper(p, v)

	assert.NoError(t, err)
	assert.Equal(t, *p.Name, "joy")
	assert.Equal(t, *p.Age, 100)
	fmt.Printf("%+v \n", p)
	fmt.Println("p.Name: ", *p.Name)
}

func TestLoadConfigFromViper_EmbedPointer(t *testing.T) {
	type Son struct {
		Name string `config:"default:ll"`
		Age  int    `config:"default:20"`
	}
	type Person struct {
		Name string `config:"default:joy"`
		Age  int    `config:"default:100"`
		Son  *Son
	}
	v := viper.New()
	var defaultValues map[string]interface{} = map[string]interface{}{
		"name":     "joy",
		"age":      100,
		"son.name": "ss",
		"son.age":  20,
	}
	SetViperDefault(v, defaultValues)

	p := &Person{}
	err := LoadConfigFromViper(p, v)

	assert.NoError(t, err)
	fmt.Printf("%+v \n", p)
	fmt.Printf("%+v \n", *p.Son)

}

//
//func TestLoadConfigFromViper_Pointer
