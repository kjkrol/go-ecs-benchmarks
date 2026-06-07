package delete10comp

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	blueprint := goke.NewBlueprint10[
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](ecs)
	var entities []goke.Entity
	for page := range blueprint.Create(n) {
		entities = append(entities, page.Entity...)
	}

	for b.Loop() {
		for _, e := range entities {
			goke.RemoveEntity(ecs, e)
		}
		b.StopTimer()

		entities = entities[:0]
		for page := range blueprint.Create(n) {
			entities = append(entities, page.Entity...)
		}
		b.StartTimer()
	}
}
