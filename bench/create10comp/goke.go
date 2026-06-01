package create10comp

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

	entities := make([]goke.Entity, 0, n)

	for range n {
		e, _, _, _, _, _, _, _, _, _, _ := blueprint.Create()
		// Just for fairness, because the others need to do that, too.
		entities = append(entities, e)
	}

	for _, e := range entities {
		goke.RemoveEntity(ecs, e)
	}
	entities = entities[:0]

	for b.Loop() {
		for range n {
			e, _, _, _, _, _, _, _, _, _, _ := blueprint.Create()
			// Just for fairness, because the others need to do that, too.
			entities = append(entities, e)
		}
		b.StopTimer()

		for _, e := range entities {
			goke.RemoveEntity(ecs, e)
		}
		entities = entities[:0]
		b.StartTimer()
	}
}
