package addremovelarge

import (
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var c1 goke.Comp[comps.C1]
	var c2 goke.Comp[comps.C2]
	var c3 goke.Comp[comps.C3]
	var c4 goke.Comp[comps.C4]
	var c5 goke.Comp[comps.C5]
	var c6 goke.Comp[comps.C6]
	var c7 goke.Comp[comps.C7]
	var c8 goke.Comp[comps.C8]
	var c9 goke.Comp[comps.C9]
	factory := ecs.NewFactory(&pos, &c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9)

	var entities []uid.UID64
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	var c10 goke.Comp[comps.C10]
	addC10 := ecs.NewEditorBuilder(&c10).Build()
	for _, e := range entities {
		addC10.Update(e)
	}

	var vel goke.Comp[comps.Velocity]
	addVel := ecs.NewEditorBuilder(&vel).Build()
	delVel := ecs.NewEditorBuilder().Delete(goke.Del[comps.Velocity]()).Build()

	for _, e := range entities {
		addVel.Update(e)
	}
	for _, e := range entities {
		delVel.Update(e)
	}

	for b.Loop() {
		for _, e := range entities {
			addVel.Update(e)
		}
		for _, e := range entities {
			delVel.Update(e)
		}
	}
}
