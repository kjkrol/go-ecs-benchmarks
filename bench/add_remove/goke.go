package addremove

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	goke.RegisterComponent[comps.Position](ecs)
	velDesc := goke.RegisterComponent[comps.Velocity](ecs)

	posBP := goke.NewBlueprint1[comps.Position](ecs)

	entities := make([]goke.Entity, 0, n)
	for range n {
		entity, _ := posBP.Create()
		entities = append(entities, entity)
	}

	for _, e := range entities {
		goke.EnsureComponent[comps.Velocity](ecs, e, velDesc)
	}
	for _, e := range entities {
		goke.RemoveComponent(ecs, e, velDesc)
	}

	for b.Loop() {
		for _, e := range entities {
			goke.EnsureComponent[comps.Velocity](ecs, e, velDesc)
		}
		for _, e := range entities {
			goke.RemoveComponent(ecs, e, velDesc)
		}
	}
}
