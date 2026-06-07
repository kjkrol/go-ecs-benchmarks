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

	realBlueprint := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)
	for page := range realBlueprint.Create(n) {
		_ = page
	}

	noiseBlueprint := goke.NewBlueprint1[comps.Position](ecs)
	i := 0
	for page := range noiseBlueprint.Create(n * 4) {
		for _, e := range page.Entity {
			if i&(1<<0) != 0 {
				goke.EnsureComponent[comps.C1](ecs, e, c1ID)
			}
			if i&(1<<1) != 0 {
				goke.EnsureComponent[comps.C2](ecs, e, c2ID)
			}
			if i&(1<<2) != 0 {
				goke.EnsureComponent[comps.C3](ecs, e, c3ID)
			}
			if i&(1<<3) != 0 {
				goke.EnsureComponent[comps.C4](ecs, e, c4ID)
			}
			if i&(1<<4) != 0 {
				goke.EnsureComponent[comps.C5](ecs, e, c5ID)
			}
			if i&(1<<5) != 0 {
				goke.EnsureComponent[comps.C6](ecs, e, c6ID)
			}
			if i&(1<<6) != 0 {
				goke.EnsureComponent[comps.C7](ecs, e, c7ID)
			}
			if i&(1<<7) != 0 {
				goke.EnsureComponent[comps.C8](ecs, e, c8ID)
			}
			i++
		}
	}

	view := goke.NewView2[comps.Position, comps.Velocity](ecs)

	loop := func() {
		for page := range view.All() {
			for i := range page.Entity {
				pos, vel := &page.Comp1[i], &page.Comp2[i]
				pos.X += vel.X
				pos.Y += vel.Y
			}
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	for page := range view.All() {
		for i := range page.Entity {
			pos := &page.Comp1[i]
			sum += pos.X + pos.Y
		}
	}
	runtime.KeepAlive(sum)
}
