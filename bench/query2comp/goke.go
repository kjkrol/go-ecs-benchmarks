package query2comp

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/kjkrol/goke/v2"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runGOKe(b *testing.B, n int) {
	ecs := goke.New()

	var pos goke.Comp[comps.Position]
	var vel goke.Comp[comps.Velocity]

	factory1 := ecs.NewFactory(&pos)
	factory1.Create(n * 10)
	for factory1.Next() {
	}

	factory2 := ecs.NewFactory(&pos, &vel)
	factory2.Create(n)
	cursor := &factory2.Cursor
	for factory2.Next() {
		velSlice := pos.Slice(cursor)
		for i := range cursor.IDs {
			velSlice[i].X, velSlice[i].Y = 1, 1
		}
	}

	query := ecs.NewQueryBuilder(&pos, &vel).Build()
	cursor = &query.Cursor
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
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
