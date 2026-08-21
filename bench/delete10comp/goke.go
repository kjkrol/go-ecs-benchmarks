package delete10comp

import (
	"testing"
	"time"

	"github.com/kjkrol/goke/v3"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var c1 goke.Comp[comps.C1]
	var c2 goke.Comp[comps.C2]
	var c3 goke.Comp[comps.C3]
	var c4 goke.Comp[comps.C4]
	var c5 goke.Comp[comps.C5]
	var c6 goke.Comp[comps.C6]
	var c7 goke.Comp[comps.C7]
	var c8 goke.Comp[comps.C8]
	var c9 goke.Comp[comps.C9]
	var c10 goke.Comp[comps.C10]
	var factory *goke.Factory
	var query *goke.Query
	var remover *goke.Remover
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		query = si.NewQueryBuilder(&c1).Build()
		remover = si.Remover()
	}})

	factory.Create(n)
	for factory.Next() {
	}

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

	for b.Loop() {
		ecs.Tick(0)
		b.StopTimer()

		factory.Create(n)
		for factory.Next() {
		}
		b.StartTimer()
	}
}
