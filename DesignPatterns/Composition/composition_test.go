package Composition

import "testing"

func TestComposition(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	localswim := Swim
	swimmer.MySwim = localswim
	swimmer.MySwim()

	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestCompositionB(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer.Train()
	swimmer.Swim()
}
