package query256arch

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
	factory := ecs.NewFactory(&pos, &vel)
	factory.Create(n)
	for factory.Next() {
	}

	noiseEntities := make([]uid.UID64, 0, n*4)
	noiseFactory := ecs.NewFactory(&pos)
	noiseFactory.Create(n * 4)
	for noiseFactory.Next() {
		noiseEntities = append(noiseEntities, noiseFactory.IDs...)
	}

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

	// Fragment the noise entities across up to 256 archetypes only after the
	// factory loop is fully drained — interleaving structural edits with an
	// active Factory iteration over the same table is unsafe.
	for i, e := range noiseEntities {
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
