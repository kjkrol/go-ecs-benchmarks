package create2comp

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()
	// GOKe blueprints always allocate memory. The only exceptions are tags, which are zero-sized structures.
	// However, there is no option to craete and have an entity with only tags components.
	blueprint := goke.NewBlueprint1[comps.Position](ecs, goke.Include[comps.Tag1]())
	entities := make([]goke.Entity, 0, n)

	for range n {
		e, _ := blueprint.Create()
		// Just for fairness, because the others need to do that, too.
		entities = append(entities, e)
	}

	for _, e := range entities {
		goke.RemoveEntity(ecs, e)
	}
	entities = entities[:0]

	for b.Loop() {
		for range n {
			e, _ := blueprint.Create()
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
