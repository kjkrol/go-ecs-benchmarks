package query256arch

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
	var c6 goke.Comp[comps.C6]
	var c7 goke.Comp[comps.C7]
	var c8 goke.Comp[comps.C8]
	var factory, noiseFactory *goke.Factory
	var query, noiseQuery *goke.Query
	var adders [8]*goke.Editor
	ecs.Setup(goke.SystemFn{OnInit: func(si *goke.SysInit) {
		factory = si.NewFactory(&pos, &vel)
		noiseFactory = si.NewFactory(&pos)
		query = si.NewQueryBuilder(&pos, &vel).Build()
		noiseQuery = si.NewQueryBuilder(&pos).Exclude(goke.Exclude[comps.Velocity]()).Build()
		adders[0] = noiseQuery.NewEditorBuilder(&c1).Build()
		adders[1] = noiseQuery.NewEditorBuilder(&c2).Build()
		adders[2] = noiseQuery.NewEditorBuilder(&c3).Build()
		adders[3] = noiseQuery.NewEditorBuilder(&c4).Build()
		adders[4] = noiseQuery.NewEditorBuilder(&c5).Build()
		adders[5] = noiseQuery.NewEditorBuilder(&c6).Build()
		adders[6] = noiseQuery.NewEditorBuilder(&c7).Build()
		adders[7] = noiseQuery.NewEditorBuilder(&c8).Build()
	}})

	factory.Create(n)
	for factory.Next() {
	}

	noiseEntities := make([]uid.UID64, 0, n*4)
	noiseFactory.Create(n * 4)
	for noiseFactory.Next() {
		noiseEntities = append(noiseEntities, noiseFactory.IDs...)
	}

	// idxByID keys a migration round by each noise entity's original
	// creation-order index (not its position in Query iteration, which
	// changes as entities migrate to new archetypes between rounds).
	idxByID := make(map[uid.UID64]int, len(noiseEntities))
	for i, e := range noiseEntities {
		idxByID[e] = i
	}

	// Fragment the noise entities across up to 256 archetypes only after the
	// factory loop is fully drained — interleaving structural edits with an
	// active Factory iteration over the same table is unsafe. noiseQuery
	// excludes Velocity so it never matches the signal (pos+vel) population.
	var fragSys [8]goke.Runnable
	for bit := range adders {
		bit, editor := bit, adders[bit]
		fragSys[bit] = ecs.RegSys(goke.SystemFn{OnUpdate: func(cb *goke.CmdBuf, _ time.Duration) {
			m := 1 << bit
			noiseQuery.All()
			for noiseQuery.Next() {
				buf := noiseQuery.BeginMigrate(cb)
				for _, id := range noiseQuery.Cursor().IDs {
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
