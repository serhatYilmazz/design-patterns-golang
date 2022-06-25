package query

type Argument int

const (
	Attack Argument = iota
	Defense
)

// Query is managing each chain's specific conditions
type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}
