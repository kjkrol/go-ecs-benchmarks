package delete2comp

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	blueprint := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)
	entities := make([]goke.Entity, 0, n)
	for range n {
		e, _, _ := blueprint.Create()
		entities = append(entities, e)
	}

	for b.Loop() {
		for _, e := range entities {
			goke.RemoveEntity(ecs, e)
		}
		b.StopTimer()

		entities = entities[:0]

		for range n {
			e, _, _ := blueprint.Create()
			entities = append(entities, e)
		}
		b.StartTimer()
	}
}
