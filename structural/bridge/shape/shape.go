package shape

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/bridge/color"
)

type Shape interface {
	Draw() string
}

type Circle struct {
	Color color.Color
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("%s%s\n\n", c.Color.Fill(), "It is a circle")
}

type Square struct {
	Color color.Color
}

func (s *Square) Draw() string {
	return fmt.Sprintf("%s%s\n\n", s.Color.Fill(), "It is a square")
}