package query256arch

import (
	"runtime"
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var realPos goke.Comp[comps.Position]
	var realVel goke.Comp[comps.Velocity]
	realFactory := ecs.NewFactory(&realPos, &realVel)
	realFactory.Create(n)
	for realFactory.Next() {
	}

	var noisePos goke.Comp[comps.Position]
	noiseFactory := ecs.NewFactory(&noisePos)

	var c1 goke.Comp[comps.C1]
	var c2 goke.Comp[comps.C2]
	var c3 goke.Comp[comps.C3]
	var c4 goke.Comp[comps.C4]
	var c5 goke.Comp[comps.C5]
	var c6 goke.Comp[comps.C6]
	var c7 goke.Comp[comps.C7]
	var c8 goke.Comp[comps.C8]
	addC1 := ecs.NewEditorBuilder(&c1).Build()
	addC2 := ecs.NewEditorBuilder(&c2).Build()
	addC3 := ecs.NewEditorBuilder(&c3).Build()
	addC4 := ecs.NewEditorBuilder(&c4).Build()
	addC5 := ecs.NewEditorBuilder(&c5).Build()
	addC6 := ecs.NewEditorBuilder(&c6).Build()
	addC7 := ecs.NewEditorBuilder(&c7).Build()
	addC8 := ecs.NewEditorBuilder(&c8).Build()

	i := 0
	noiseFactory.Create(n * 4)
	for noiseFactory.Next() {
		for _, e := range noiseFactory.IDs {
			if i&(1<<0) != 0 {
				addC1.Update(e)
			}
			if i&(1<<1) != 0 {
				addC2.Update(e)
			}
			if i&(1<<2) != 0 {
				addC3.Update(e)
			}
			if i&(1<<3) != 0 {
				addC4.Update(e)
			}
			if i&(1<<4) != 0 {
				addC5.Update(e)
			}
			if i&(1<<5) != 0 {
				addC6.Update(e)
			}
			if i&(1<<6) != 0 {
				addC7.Update(e)
			}
			if i&(1<<7) != 0 {
				addC8.Update(e)
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
