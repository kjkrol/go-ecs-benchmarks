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

	for range n * 10 {
		_, _ = posBP.Create()
	}
	for range n {
		_, _, v := posVelBP.Create()
		v.X, v.Y = 1, 1
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
		pos, _ := head.V1, head.V2
		sum += pos.X + pos.Y
	}
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
