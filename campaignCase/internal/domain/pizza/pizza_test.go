package pizza

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePizza(t *testing.T) {
    assert := assert.New(t)

    name := "Peperone"
    size := 2

    sauces := []string{
        "Elefant",
        "Arisco",
    }

    ingredients := Ingredients{
        Sauces: sauces,
    }

    pizza := CreatePizza(name, size, ingredients)

    assert.Equal(pizza.Name, "name")
}