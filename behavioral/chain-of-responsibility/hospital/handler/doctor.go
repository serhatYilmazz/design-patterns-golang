package handler

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/cOr/patient"
)

type DoctorExamination struct {
	next Department
}

func (d *DoctorExamination) Execute(patient *patient.Patient) {
	if !patient.IsExamined {
		fmt.Println("Patient is examining by a doctor now")
		patient.IsExamined = true

		d.next.Execute(patient)
		return
	}

	fmt.Println("Patient is already examined")
	d.next.Execute(patient)
}

func (d *DoctorExamination) SetNext(department Department) Department {
	d.next = department
	return department
}
