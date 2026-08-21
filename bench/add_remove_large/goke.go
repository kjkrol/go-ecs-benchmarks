package addremovelarge

import (
	"testing"
	"time"

	"github.com/kjkrol/goke/v3"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]
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
	var withoutVel, withVel *goke.Query
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&pos, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)
		withoutVel = si.NewQueryBuilder(&pos).Exclude(goke.Exclude[comps.Velocity]()).Build()
		withVel = si.NewQueryBuilder(&pos).Include(goke.Include[comps.Velocity]()).Build()
	}})

	var entities []uid.UID64
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	addVel := withoutVel.NewEditorBuilder(&vel).Build()
	delVel := withVel.NewEditorBuilder().Remove(goke.Remove[comps.Velocity]()).Build()

	addSys := ecs.RegSys(goke.SystemFn{OnUpdate: func(cb *goke.CmdBuf, _ time.Duration) {
		withoutVel.All()
		for withoutVel.Next() {
			buf := withoutVel.BeginMigrate(cb)
			for _, id := range withoutVel.Cursor().IDs {
				buf.Add(id)
			}
			buf.Commit(addVel)
		}
	}})
	delSys := ecs.RegSys(goke.SystemFn{OnUpdate: func(cb *goke.CmdBuf, _ time.Duration) {
		withVel.All()
		for withVel.Next() {
			buf := withVel.BeginMigrate(cb)
			for _, id := range withVel.Cursor().IDs {
				buf.Add(id)
			}
			buf.Commit(delVel)
		}
	}})
	ecs.SetPlan(func(ctx goke.RunCtx, d time.Duration) {
		ctx.Run(addSys, d)
		_ = ctx.Sync()
		ctx.Run(delSys, d)
		_ = ctx.Sync()
	})

	ecs.Tick(0)
	for b.Loop() {
		ecs.Tick(0)
	}
}
