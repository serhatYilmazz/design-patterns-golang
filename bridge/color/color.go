package color

import "fmt"

type Color interface {
	Fill() string
}

type Red struct {}

func (r *Red) Fill() string {
	return fmt.Sprintf("Color is red\n")
}

type Green struct {}

func (g *Green) Fill() string {
	return fmt.Sprintf("Color is green\n")
}
