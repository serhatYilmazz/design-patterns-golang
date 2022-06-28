package handler

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/cOr/patient"
)

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(patient *patient.Patient) {
	if !patient.IsGoCashier {
		fmt.Println("Patient is paying the price")
		patient.IsGoCashier = true
		return
	}

	fmt.Println("Patient is already paid the price")
}

func (c *Cashier) SetNext(department Department) Department {
	c.next = department
	return department
}
