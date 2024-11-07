//source gpt
package builder

import "fmt"

// Car represents the complex object we want to build
type Car struct {
	Make       string
	Model      string
	Color      string
	EngineType string
	Wheels     int
}

// CarBuilder is the builder struct that constructs Car objects
type CarBuilder struct {
	make       string
	model      string
	color      string
	engineType string
	wheels     int
}

// NewCarBuilder creates a new instance of CarBuilder
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

// SetMake sets the make of the car
func (b *CarBuilder) SetMake(make string) *CarBuilder {
	b.make = make
	return b
}

// SetModel sets the model of the car
func (b *CarBuilder) SetModel(model string) *CarBuilder {
	b.model = model
	return b
}

// SetColor sets the color of the car
func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.color = color
	return b
}

// SetEngineType sets the engine type of the car
func (b *CarBuilder) SetEngineType(engineType string) *CarBuilder {
	b.engineType = engineType
	return b
}

// SetWheels sets the number of wheels of the car
func (b *CarBuilder) SetWheels(wheels int) *CarBuilder {
	b.wheels = wheels
	return b
}

// Build constructs and returns the final Car object
func (b *CarBuilder) Build() *Car {
	return &Car{
		Make:       b.make,
		Model:      b.model,
		Color:      b.color,
		EngineType: b.engineType,
		Wheels:     b.wheels,
	}
}

func main2() {
	// Using the builder to create a Car instance step-by-step
	car := NewCarBuilder().
		SetMake("Toyota").
		SetModel("Camry").
		SetColor("Red").
		SetEngineType("Hybrid").
		SetWheels(4).
		Build()

	fmt.Printf("Car built: %+v\n", car)
}