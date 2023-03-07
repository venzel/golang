package auto

type Auto struct {
    Name string
    Year int
}

type Car struct {
    Auto
    Color string
}

type Truck struct {
    Auto
    Wheels int
}

func (car *Car) SetColor(color string) {
    car.Color = color;
}

func (car *Car) GetColor() string {
    return car.Color;
}