package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not	recognized\n", f))
	}
}
