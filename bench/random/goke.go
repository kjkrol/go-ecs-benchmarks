package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]

	entities := make([]uid.UID64, 0, n)

	factory := ecs.NewFactory(&pos)
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	// Don't use b.Loop and callback, as we do not want to measure
	// the cost of calling the non-inlined callback.
	b.ResetTimer()
	query := ecs.NewQueryBuilder(&pos).Build()
	cursor := &query.Cursor
	for range b.N {
		for _, e := range entities {
			query.Seek(e)
			pos := pos.At(cursor)
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
