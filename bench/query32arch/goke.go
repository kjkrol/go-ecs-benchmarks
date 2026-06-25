package query32arch

import (
	"runtime"
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]

	entities := make([]uid.UID64, 0, n)

	factory := ecs.NewFactory(&pos, &vel)
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	var c1 goke.Comp[comps.C1]
	var c2 goke.Comp[comps.C2]
	var c3 goke.Comp[comps.C3]
	var c4 goke.Comp[comps.C4]
	var c5 goke.Comp[comps.C5]
	adders := [5]*goke.Editor{
		ecs.NewEditorBuilder(&c1).Build(),
		ecs.NewEditorBuilder(&c2).Build(),
		ecs.NewEditorBuilder(&c3).Build(),
		ecs.NewEditorBuilder(&c4).Build(),
		ecs.NewEditorBuilder(&c5).Build(),
	}

	// Distribute entities across up to 32 archetypes only after the factory
	// loop is fully drained — interleaving structural edits with an active
	// Factory iteration over the same table is unsafe.
	for i, e := range entities {
		for j, ed := range adders {
			m := 1 << j
			if i&m == m {
				ed.Update(e)
			}
		}
	}

	query := ecs.NewQueryBuilder(&pos, &vel).Build()
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
