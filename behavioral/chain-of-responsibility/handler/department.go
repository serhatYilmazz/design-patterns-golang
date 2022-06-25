package handler

import "github.com/serhatYilmazz/design-patterns-golang/behavioral/cOr/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department) Department
}
