package main

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/bridge/color"
	"github.com/serhatYilmazz/design-patterns-golang/bridge/shape"
)

func main() {
	red := &color.Red{}
	green := &color.Green{}

	redCircle := shape.Circle{Color: red}
	greenCircle := shape.Circle{Color: green}
	redSquare := shape.Square{Color: red}
	greenSquare := shape.Square{Color: green}

	fmt.Println(redCircle.Draw())
	fmt.Println(greenCircle.Draw())
	fmt.Println(redSquare.Draw())
	fmt.Println(greenSquare.Draw())
}
