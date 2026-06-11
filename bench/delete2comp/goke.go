package delete2comp

import (
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]
	factory := ecs.NewFactory(&pos, &vel)

	entities := make([]uid.UID64, 0, n)
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	for range b.N {
		for _, e := range entities {
			ecs.RemoveEnt(e)
		}
		b.StopTimer()

		entities = entities[:0]

		factory.Create(n)
		for factory.Next() {
			entities = append(entities, factory.IDs...)
		}
		b.StartTimer()
	}
}
