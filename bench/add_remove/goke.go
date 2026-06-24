package addremove

import (
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]

	posBP := ecs.NewFactory(&pos)

	var entities []uid.UID64
	posBP.Create(n)
	for posBP.Next() {
		entities = append(entities, posBP.IDs...)
	}

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
