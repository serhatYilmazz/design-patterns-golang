package mediator

import (
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/observer"
	"github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/query"
	"sync"
)

// Game is the mediator of the system
type Game struct {
	Observers sync.Map
}

func (g *Game) Subscribe(o observer.Observer) {
	g.Observers.Store(o, struct{}{})
}

func (g *Game) UnSubscribe(o observer.Observer) {
	g.Observers.Delete(o)
}

func (g *Game) Fire(q *query.Query) {
	g.Observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(observer.Observer).Handle(q)
		return true
	})
}
