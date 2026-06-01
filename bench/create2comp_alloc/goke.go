package create2compalloc

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	for b.Loop() {
		b.StopTimer()
		ecs := goke.New()

		blueprint := goke.NewBlueprint2[comps.Position, comps.Velocity](ecs)

		b.StartTimer()
		for range n {
			_, p, v := blueprint.Create()
			p.X, p.Y = 1, 1
			v.X, v.Y = 1, 1
		}
	}
}
