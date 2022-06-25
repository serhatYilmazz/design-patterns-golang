package main

import (
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/cOr/handler"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/cOr/patient"
)

func main() {
	cashier := handler.Cashier{}
	examination := handler.DoctorExamination{}
	reception := handler.Reception{}

	reception.SetNext(&examination).SetNext(&cashier)

	p := patient.Patient{}

	reception.Execute(&p)
}
