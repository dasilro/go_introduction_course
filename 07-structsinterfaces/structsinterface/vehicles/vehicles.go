package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
	GokuVehicle       = "GOKU"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	case GokuVehicle:
		return &Goku{Time: time}, nil
	}

	return nil, fmt.Errorf("vehicle '%s' doesn't exist", v)
}

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

type Motorcycle struct {
	Time int
}

func (m *Motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60)
}

type Truck struct {
	Time int
}

func (t *Truck) Distance() float64 {
	return 70 * (float64(t.Time) / 60)
}

type Goku struct {
	Time int
}

func (g *Goku) Distance() float64 {
	return 300000 * (float64(g.Time) / 60)
}
