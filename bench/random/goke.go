package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/kjkrol/goke/v3"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var factory *goke.Factory
	var query *goke.Query
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&pos)
		query = si.NewQueryBuilder(&pos).Build()
	}})

	entities := make([]uid.UID64, 0, n)
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	// Don't use b.Loop and callback, as we do not want to measure
	// the cost of calling the non-inlined callback.
	b.ResetTimer()
	cursor := query.Cursor()
	// All entities come from a single Factory.Create call, so they share one
	// archetype: Seek once to establish it, then SeekH for the rest.
	query.Seek(entities[0])
	for range b.N {
		for _, e := range entities {
			query.SeekH(e)
			pos := pos.At(cursor)
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
