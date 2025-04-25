package gofaker

import (
	"math/rand"
	"time"
)

type Faker struct {
	Rand *rand.Rand
}

func New(seed uint64) *Faker {
	var source rand.Source
	if seed == 0 {
		source = rand.NewSource(time.Now().UnixNano())
	} else {
		source = rand.NewSource(int64(seed))
	}

	return &Faker{
		Rand: rand.New(source),
	}
}
