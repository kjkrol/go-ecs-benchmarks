package delete10comp

import (
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/kjkrol/uid"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

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
	factory := ecs.NewFactory(&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10)

	var entities []uid.UID64
	factory.Create(n)
	for factory.Next() {
		entities = append(entities, factory.IDs...)
	}

	for range b.N {
		for _, e := range entities {
			ecs.RemoveEnt(e)
		}
		b.StopTimer()

		entities = entities[:0]
		factory.Create(n)
		for factory.Next() {
			entities = append(entities, factory.IDs...)
		}
		b.StartTimer()
	}
}
