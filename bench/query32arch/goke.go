package query32arch

import (
	"runtime"
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
	var factory *goke.Factory
	var query *goke.Query
	var adders [5]*goke.Editor
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&pos, &vel)
		query = si.NewQueryBuilder(&pos, &vel).Build()
		adders[0] = query.NewEditorBuilder(&c1).Build()
		adders[1] = query.NewEditorBuilder(&c2).Build()
		adders[2] = query.NewEditorBuilder(&c3).Build()
		adders[3] = query.NewEditorBuilder(&c4).Build()
		adders[4] = query.NewEditorBuilder(&c5).Build()
	}})

	entities := make([]uid.UID64, 0, n)
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	// idxByID keys a migration round by each entity's original creation-order
	// index (not its position in Query iteration, which changes as entities
	// migrate to new archetypes between rounds).
	idxByID := make(map[uid.UID64]int, n)
	for i, e := range entities {
		idxByID[e] = i
	}

	// Distribute entities across up to 32 archetypes only after the factory
	// loop is fully drained — interleaving structural edits with an active
	// Factory iteration over the same table is unsafe.
	var fragSys [5]goke.Runnable
	for bit := range adders {
		bit, editor := bit, adders[bit]
		fragSys[bit] = ecs.RegSys(goke.SystemFn{OnUpdate: func(cb *goke.CmdBuf, _ time.Duration) {
			m := 1 << bit
			query.All()
			for query.Next() {
				buf := query.BeginMigrate(cb)
				for _, id := range query.Cursor().IDs {
					if idxByID[id]&m == m {
						buf.Add(id)
					}
				}
				buf.Commit(editor)
			}
		}})
	}
	ecs.SetPlan(func(ctx goke.RunCtx, d time.Duration) {
		for _, sys := range fragSys {
			ctx.Run(sys, d)
			_ = ctx.Sync()
		}
	})
	ecs.Tick(0)

	cursor := query.Cursor()
	loop := func() {
		query.All()
		for query.Next() {
			posSlice := pos.Slice(cursor)
			velSlice := vel.Slice(cursor)
			for i := range cursor.IDs {
				posSlice[i].X += velSlice[i].X
				posSlice[i].Y += velSlice[i].Y
			}
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query.All()
	for query.Next() {
		posSlice := pos.Slice(cursor)
		for i := range cursor.IDs {
			sum += posSlice[i].X + posSlice[i].Y
		}
	}
	runtime.KeepAlive(sum)
}
