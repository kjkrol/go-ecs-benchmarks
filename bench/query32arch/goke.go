package query32arch

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

	extraIDs := []goke.ComponentDesc{c1ID, c2ID, c3ID, c4ID, c5ID}
	blueprint := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)

	i := 0
	for page := range blueprint.Create(n) {
		for _, e := range page.Entity {
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
					default:
						panic("Unknown component type: " + id.Type.String())
					}
				}
			}
			i++
		}
	}

	view := goke.NewView2[comps.Position, comps.Velocity](ecs)

	loop := func() {
		for page := range view.All() {
			for i, _ := range page.Entity {
				page.Comp1[i].X += page.Comp2[i].X
				page.Comp1[i].Y += page.Comp2[i].Y
			}
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	for head := range view.All() {
		for i, _ := range head.Entity {
			sum += head.Comp1[i].X + head.Comp1[i].Y
		}
	}
	runtime.KeepAlive(sum)
}
