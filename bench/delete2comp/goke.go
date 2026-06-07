package delete2comp

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

	for b.Loop() {
		for _, e := range entities {
			goke.RemoveEntity(ecs, e)
		}
		b.StopTimer()

		entities = entities[:0]

		for page := range blueprint.Create(n) {
			for _, e := range page.Entity {
				entities = append(entities, e)
			}
		}
		b.StartTimer()
	}
}
