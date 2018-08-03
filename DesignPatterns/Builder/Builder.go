package creational

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetStructure().SetWheels()
}
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
type CarBuilder struct{ v VehicleProduct }

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

type BikeBuilder struct{ v VehicleProduct }

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}
func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}
func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}
func (c *BikeBuilder) Build() VehicleProduct {
	return c.v
}
