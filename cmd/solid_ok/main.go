package main

import "fmt"

// Tool - інтерфейс для стоматологічних інструментів, дотримуючись принципу Interface Segregation
type Tool interface {
	Use()
}

// Drill - клас для свердла, що дотримується принципу Single Responsibility
type Drill struct {
	power int
}

func (d *Drill) Use() {
	fmt.Println("Using the drill at power:", d.power)
}

func NewDrill(power int) *Drill {
	return &Drill{power: power}
}

// Scaler - клас для скалера, що дотримується принципу Single Responsibility
type Scaler struct {
	frequency int
}

func (s *Scaler) Use() {
	fmt.Println("Using the scaler at frequency:", s.frequency)
}

func NewScaler(frequency int) *Scaler {
	return &Scaler{frequency: frequency}
}

// DentalStation - клас для стоматологічного стілу, дотримуючись принципу Open/Closed та Dependency Inversion
type DentalStation struct {
	tools []Tool
}

func (d *DentalStation) AddTool(tool Tool) {
	d.tools = append(d.tools, tool)
}

func (d *DentalStation) UseAllTools() {
	for _, tool := range d.tools {
		tool.Use()
	}
}

func main() {
	drill := NewDrill(5)
	scaler := NewScaler(10)

	dentalStation := DentalStation{}
	dentalStation.AddTool(drill)
	dentalStation.AddTool(scaler)

	dentalStation.UseAllTools()

	workWithTooth(drill)
	workWithTooth(scaler)
}

func workWithTooth(t Tool) {
	// Liskov Substitution Principle виконується, тому що всі інструменти виконують операцію саме із зубами пацієнта
	t.Use()
}
