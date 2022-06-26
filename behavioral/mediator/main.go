package main

import (
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/mediator/train"
)

func main() {
	sm := train.StationManager{
		IsStationEmpty: true,
		Trains:         []train.Train{},
	}

	passengerTrain := train.PassengerTrain{M: &sm}
	freightTrain := train.FreightTrain{M: &sm}

	freightTrain.Arrive()
	passengerTrain.Arrive()
	freightTrain.Depart()
}
