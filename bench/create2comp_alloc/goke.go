package create2compalloc

import (
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {

	for range b.N {
		b.StopTimer()
		ecs := goke.New()

		var pos goke.Comp[comps.Position]
		var vel goke.Comp[comps.Velocity]
		factory := ecs.NewFactory(&pos, &vel)

		b.StartTimer()
		factory.Create(n)
		for factory.Next() {
		}
	}
}
