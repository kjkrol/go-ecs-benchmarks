package addremove

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
	var posBP *goke.Factory
	var withoutVel, withVel *goke.Query
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		posBP = si.NewFactory(&pos)
		withoutVel = si.NewQueryBuilder(&pos).Exclude(goke.Exclude[comps.Velocity]()).Build()
		withVel = si.NewQueryBuilder(&pos).Include(goke.Include[comps.Velocity]()).Build()
	}})

	var entities []uid.UID64
	posBP.Create(n)
	for posBP.Next() {
		entities = append(entities, posBP.IDs...)
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
