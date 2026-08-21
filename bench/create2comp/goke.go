package create2comp

import (
	"testing"
	"time"

	"github.com/kjkrol/goke/v3"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]
	var factory *goke.Factory
	var query *goke.Query
	var remover *goke.Remover
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&pos, &vel)
		query = si.NewQueryBuilder(&pos, &vel).Build()
		remover = si.Remover()
	}})

	// removeSys is the untimed cleanup between iterations, never the
	// measured operation.
	removeSys := ecs.RegSys(goke.SystemFn{OnUpdate: func(cb *goke.CmdBuf, _ time.Duration) {
		query.All()
		for query.Next() {
			buf := query.BeginMigrate(cb)
			for _, id := range query.Cursor().IDs {
				buf.Add(id)
			}
			buf.Commit(remover)
		}
	}})
	ecs.SetPlan(func(ctx goke.RunCtx, d time.Duration) {
		ctx.Run(removeSys, d)
		_ = ctx.Sync()
	})

	factory.Create(n)
	for factory.Next() {
	}

	ecs.Tick(0)

	for b.Loop() {
		factory.Create(n)
		for factory.Next() {
		}
		b.StopTimer()
		ecs.Tick(0)
		b.StartTimer()
	}
}
