package creature

import (
	"fmt"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/mediator"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/query"
)

type Creature struct {
	Game    *mediator.Game
	Name    string
	At, Def int
}

func NewCreature(game *mediator.Game, name string, attack int, defense int) *Creature {
	return &Creature{
		Game: game,
		Name: name,
		At:   attack,
		Def:  defense,
	}
}

func (c *Creature) Attack() int {
	q := query.Query{
		CreatureName: c.Name,
		WhatToQuery:  query.Attack,
		Value:        c.At,
	}
	c.Game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := query.Query{
		CreatureName: c.Name,
		WhatToQuery:  query.Defense,
		Value:        c.Def,
	}
	c.Game.Fire(&q)
	return q.Value
}

func (c *Creature) ToString() string {
	return fmt.Sprintf("%s (%d|%d)", c.Name, c.Attack(), c.Defense())
}

// Modifier is just a template
type Modifier struct {
	Game     *mediator.Game
	Creature *Creature
}

func (m Modifier) Handle(q *query.Query) {
	fmt.Println("It is modifier")
}

type DoubleAttackModifier struct {
	Modifier
}

func NewDoubleAttackModifier(g *mediator.Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{Modifier{
		Game:     g,
		Creature: c,
	}}
	g.Subscribe(d) // OOP critical part, When DoubleAttackModifier has named fields and not has Handle() method, it is not getting compiled.
	return d
}

func (d *DoubleAttackModifier) Handle(q *query.Query) {
	if q.CreatureName == d.Creature.Name && q.WhatToQuery == query.Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.Game.UnSubscribe(d)
	return nil
}
