package query32arch

import (
	"runtime"
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]
	factory := ecs.NewFactory(&pos, &vel)

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

	i := 0
	factory.Create(n)
	for factory.Next() {
		for _, e := range factory.IDs {
			for j, ed := range adders {
				m := 1 << j
				if i&m == m {
					ed.Update(e)
				}
			}
			i++
		}
	}

	var qpos goke.Comp[comps.Position]
	var qvel goke.Comp[comps.Velocity]
	query := ecs.NewQueryBuilder(&qpos, &qvel).Build()

	loop := func() {
		query.All()
		for query.Next() {
			posSlice := qpos.Slice(&query.Cursor)
			velSlice := qvel.Slice(&query.Cursor)
			for i := range posSlice {
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
		posSlice := qpos.Slice(&query.Cursor)
		for i := range posSlice {
			sum += posSlice[i].X + posSlice[i].Y
		}
	}
	runtime.KeepAlive(sum)
}
