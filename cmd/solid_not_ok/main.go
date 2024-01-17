package main

import "fmt"

type DentalStation interface {
	UseScaler()
	UseDrill()
	UseLight(intensity int)
	// неможливо додати до станції новий інструмент, не переписавши код

	ChangeDrillPower(power int)
	UpdateScaler(frequency int)
	AdjustLight(intensity int)
	// Забагато методів в інтерфейсі
}

// DefaultDentalStation - один клас, який виконує багато завдань
type DefaultDentalStation struct {
	drillPower      int
	scalerFrequency int
	lightIntensity  int
}

func (d *DefaultDentalStation) ChangeDrillPower(power int) {
	d.drillPower = power
	// Зміна інших налаштувань, які не пов'язані безпосередньо з свердлом
	d.scalerFrequency += power
}

func (d *DefaultDentalStation) UpdateScaler(frequency int) {
	d.scalerFrequency = frequency
}

func (d *DefaultDentalStation) AdjustLight(intensity int) {
	d.lightIntensity = intensity
	// Логіка, яка не повинна бути частиною цієї функції
	fmt.Println("Adjusting scaler frequency along with light intensity")
}

func (d *DefaultDentalStation) UseScaler() {
	fmt.Println("Scaler is using")
}

func (d *DefaultDentalStation) UseDrill() {
	fmt.Println("Drill is using")
}

func (d *DefaultDentalStation) UseLight(intensity int) {
	d.lightIntensity = intensity
	fmt.Println(`Light is using`)
	// Логіка, яка не повинна бути частиною цієї функції
}

// DentalTool - загальний інтерфейс для всіх інструментів
type DentalTool interface {
	UseTool()
	MaintainTool()
	AdjustSettings()
	AttachToStation(*DefaultDentalStation)
}

// Специфічний інструмент, який використовує загальний інтерфейс
type Drill struct{}

func (d Drill) UseTool() {
	fmt.Println("Drill is using")
}

func (d Drill) MaintainTool() {
	fmt.Println("Maintaining the drill")
}

func (d Drill) AdjustSettings() {
	// Цей метод лише для імплементації інттерфейса
}

func (d Drill) AttachToStation(s *DefaultDentalStation) {
	fmt.Println("Drill is attaching to station")
}

type Light struct{}

func (d Light) UseTool() {
	fmt.Println("Light is shining")
}

func (d Light) MaintainTool() {
	// Цей метод лише для імплементації інттерфейса
}

func (d Light) AdjustSettings() {
	// Цей метод лише для імплементації інттерфейса
}

func (d Light) AttachToStation(s *DefaultDentalStation) {
	fmt.Println("Light is attaching to station")
}

func main() {
	var dentalStation DentalStation

	dentalStation = &DefaultDentalStation{}
	dentalStation.ChangeDrillPower(5)
	dentalStation.UpdateScaler(10)
	dentalStation.AdjustLight(7)

	var drill DentalTool

	drill = Drill{}
	drill.UseTool()
	drill.MaintainTool()
	drill.AdjustSettings()

	workWithTooth(drill)
	workWithTooth(Light{})

	//drill.AttachToStation(dentalStation)
	ds := &DefaultDentalStation{}
	drill.AttachToStation(ds)
}

func workWithTooth(t DentalTool) {
	// Liskov Substitution Principle не виконується, тому що світло не виконує жодних операцій із зубами
	t.UseTool()
}
