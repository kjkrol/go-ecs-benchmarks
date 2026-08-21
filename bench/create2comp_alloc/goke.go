package create2compalloc

import (
	"testing"

	"github.com/kjkrol/goke/v3"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {

	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]
	var factory *goke.Factory
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&pos, &vel)
	}})

	for b.Loop() {
		factory.Create(n)
		for factory.Next() {
		}
	}
}
