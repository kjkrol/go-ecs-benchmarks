package create2comp

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()
	blueprint := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)

	var entities []goke.Entity
	for page := range blueprint.Create(n) {
		for _, e := range page.Entity {
			entities = append(entities, e)
		}
	}

	for _, e := range entities {
		goke.RemoveEntity(ecs, e)
	}
	entities = entities[:0]

	for b.Loop() {
		for page := range blueprint.Create(n) {
			for _, e := range page.Entity {
				entities = append(entities, e)
			}
		}
		b.StopTimer()

		for _, e := range entities {
			goke.RemoveEntity(ecs, e)
		}
		entities = entities[:0]
		b.StartTimer()
	}
}
