package pizza

type Ingredients struct {
    Sauces []string
}

type Pizza struct {
    Name string
    Size int
    Ingredients
}

func CreatePizza(name string, size int, ingredients Ingredients) *Pizza {
    return &Pizza{
        Name: name,
        Size: size,
        Ingredients: ingredients,
    }
}

