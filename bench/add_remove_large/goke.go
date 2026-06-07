package addremovelarge

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	goke.RegisterComponent[comps.Position](ecs)
	velDesc := goke.RegisterComponent[comps.Velocity](ecs)
	c10Desc := goke.RegisterComponent[comps.C10](ecs)

	blueprint := goke.NewBlueprint10[
		comps.Position,
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9,
	](ecs)

	var entities []goke.Entity
	for page := range blueprint.Create(n) {
		for _, e := range page.Entity {
			entities = append(entities, e)
		}
	}
	for _, entity := range entities {
		goke.EnsureComponent[comps.C10](ecs, entity, c10Desc)
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
