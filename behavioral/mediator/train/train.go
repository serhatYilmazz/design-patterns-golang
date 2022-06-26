package train

import (
	"fmt"
)

type Mediator interface {
	CanArrive(t Train) bool
	NotifyAboutDeparture()
}

type StationManager struct {
	IsStationEmpty bool
	Trains         []Train
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.IsStationEmpty {
		s.IsStationEmpty = false
		return true
	}
	fmt.Println("Station is full")
	s.Trains = append(s.Trains, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.IsStationEmpty {
		s.IsStationEmpty = true
	}
	if len(s.Trains) > 0 {
		firstTrain := s.Trains[0]
		s.Trains = s.Trains[1:]
		firstTrain.PermitArrival()
	}
}

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

type PassengerTrain struct {
	M *StationManager
}

func (p *PassengerTrain) Arrive() {
	if p.M.CanArrive(p) {
		fmt.Println("Passenger train is arrived")
		return
	}

	fmt.Println("Passenger train can not arrive since there is no space in station")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("Passenger train has been departed")
	p.M.NotifyAboutDeparture()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("Passenger train arrival is permitted")
	p.Arrive()
}

type FreightTrain struct {
	M *StationManager
}

func (f *FreightTrain) Arrive() {
	if f.M.CanArrive(f) {
		fmt.Println("Freight train is arrived")
		return
	}
	fmt.Println("Passenger train can not arrive since there is no space in station")
}

func (f *FreightTrain) Depart() {
	fmt.Println("Freight train has been departed")
	f.M.NotifyAboutDeparture()
}

func (f *FreightTrain) PermitArrival() {
	fmt.Println("Freight train is allowed to arrive")
	f.Arrive()
}
