package observer

import "github.com/serhatYilmazz/design-patterns-golang/behavioral/command-query-separation/query"

// Observer is the part of the chain (of responsibility)
type Observer interface {
	Handle(q *query.Query)
}

// Observable is the main Mediator
type Observable interface {
	Subscribe(o Observer)
	UnSubscribe(o Observer)
	Fire(q *query.Query)
}
