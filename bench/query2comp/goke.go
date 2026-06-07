package query2comp

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	goke.RegisterComponent[comps.Position](ecs)
	goke.RegisterComponent[comps.Velocity](ecs)

	posBP := goke.NewBlueprint1[comps.Position](ecs)
	posVelBP := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)

	for _ = range posBP.Create(n * 10) {
	}

	for page := range posVelBP.Create(n) {
		for i, _ := range page.Entity {
			v := &page.Comp2[i]
			v.X, v.Y = 1, 1
		}
	}

	view := goke.NewView2[comps.Position, comps.Velocity](ecs)

	loop := func() {
		for page := range view.All() {
			for i, _ := range page.Entity {
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
		for i, _ := range page.Entity {
			pos, vel := &page.Comp1[i], &page.Comp2[i]
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
