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
	factory := ecs.NewFactory(&pos)

	entities := make([]uid.UID64, 0, n)
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}
	rand.Shuffle(n, util.Swap(entities))

	var qpos goke.Comp[comps.Position]
	query := ecs.NewQueryBuilder(&qpos).Build()

	sum := 0.0
	// Don't use b.Loop and callback, as we do not want to measure
	// the cost of calling the non-inlined callback.
	b.ResetTimer()
	for range b.N {
		query.Pick(entities)
		for query.Next() {
			pos := qpos.At(&query.Cursor)
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
