package query256arch

import (
	"runtime"
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	goke.RegisterComponent[comps.Position](ecs)
	goke.RegisterComponent[comps.Velocity](ecs)
	c1ID := goke.RegisterComponent[comps.C1](ecs)
	c2ID := goke.RegisterComponent[comps.C2](ecs)
	c3ID := goke.RegisterComponent[comps.C3](ecs)
	c4ID := goke.RegisterComponent[comps.C4](ecs)
	c5ID := goke.RegisterComponent[comps.C5](ecs)
	c6ID := goke.RegisterComponent[comps.C6](ecs)
	c7ID := goke.RegisterComponent[comps.C7](ecs)
	c8ID := goke.RegisterComponent[comps.C8](ecs)

	extraIDs := []goke.ComponentDesc{c1ID, c2ID, c3ID, c4ID, c5ID, c6ID, c7ID, c8ID}

	blueprint := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)

	for i := range n * 4 {
		e, _, _ := blueprint.Create()
		for j, id := range extraIDs {
			m := 1 << j
			if i&m == m {
				switch id {
				case c1ID:
					goke.EnsureComponent[comps.C1](ecs, e, id)
				case c2ID:
					goke.EnsureComponent[comps.C2](ecs, e, id)
				case c3ID:
					goke.EnsureComponent[comps.C3](ecs, e, id)
				case c4ID:
					goke.EnsureComponent[comps.C4](ecs, e, id)
				case c5ID:
					goke.EnsureComponent[comps.C5](ecs, e, id)
				case c6ID:
					goke.EnsureComponent[comps.C6](ecs, e, id)
				case c7ID:
					goke.EnsureComponent[comps.C7](ecs, e, id)
				case c8ID:
					goke.EnsureComponent[comps.C8](ecs, e, id)
				default:
					panic("Unknown component type: " + id.Type.String())
				}
			}
		}
	}

	view := goke.NewView2[comps.Position, comps.Velocity](ecs)

	loop := func() {
		for head := range view.Values() {
			pos, vel := head.V1, head.V2
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	for head := range view.Values() {
		pos := head.V1
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}
