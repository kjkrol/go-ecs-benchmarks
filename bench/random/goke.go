package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/kjkrol/goke"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	goke.RegisterComponent[comps.Position](ecs)

	blueprint := goke.NewBlueprint1[comps.Position](ecs)
	view := goke.NewView1[comps.Position](ecs)

	entities := make([]goke.Entity, 0, n)
	for page := range blueprint.Create(n) {
		entities = append(entities, page.Entity...)
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	// Don't use b.Loop and callback, as we do not want to measure
	// the cost of calling the non-inlined callback.
	b.ResetTimer()
	for range b.N {
		for _, item := range view.Filter(entities) {
			pos := item.Comp1
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
