package main

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/creature"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/mediator"
	"sync"
)

func main() {
	game := mediator.Game{Observers: sync.Map{}} // Central Mediator

	goblin := creature.NewCreature(&game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.ToString())

	m := creature.NewDoubleAttackModifier(&game, goblin)
	fmt.Println(goblin.ToString())
	m.Close()

	fmt.Println(goblin.ToString())
}
