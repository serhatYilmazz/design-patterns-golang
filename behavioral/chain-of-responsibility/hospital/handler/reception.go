package handler

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/cOr/patient"
)

type Reception struct {
	next Department
}

func (r *Reception) SetNext(nextDep Department) Department {
	r.next = nextDep
	return nextDep
}

func (r *Reception) Execute(p *patient.Patient) {
	if !p.IsReserved {
		fmt.Println("Patient has reservation now")
		p.IsReserved = true

		r.next.Execute(p)
		return
	}
	fmt.Println("Patient is already reserved")
	r.next.Execute(p)
}
