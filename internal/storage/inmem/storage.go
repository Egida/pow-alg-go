package inmem

import (
	"context"
	"math/rand"
)

type Storage struct {
	quotes []string
}

func NewStorage() *Storage {
	return &Storage{quotes: quotes}
}

func (s *Storage) GetRandomQuote(_ context.Context) (string, error) {
	idx := rand.Intn(len(s.quotes))
	return s.quotes[idx], nil
}

var quotes = []string{
	"Wisdom is but a powerful tree in the mind's forest. It grows stronger, sturdier, and with time becomes an undeniable component of a person's landscape.",
	"Learn to listen, and you will grow wise.",
	"Wisdom cannot be rushed. It comes in its own time.",
	"Patience and time grow the seed that blossoms into wisdom.",
	"Wisdom lives inside all who are quiet enough to listen for it.",
	"With every wrinkle is a story, a lesson, or a step in wisdom and realization.",
	"Growing older means being wise enough to know when to let things go and let things be.",
	"Never decide you are smart enough. Be wise enough to recognize that there is always more to learn.",
	"A wise man never wears his wisdom like a badge.",
	"Growing wise means learning what to put your energy into and what to walk away from.",
}
