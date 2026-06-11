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
	posBP := ecs.NewFactory(&pos)

	var posVelPos goke.Comp[comps.Position]
	var posVelVel goke.Comp[comps.Velocity]
	posVelBP := ecs.NewFactory(&posVelPos, &posVelVel)

	posBP.Create(n * 10)
	for posBP.Next() {
	}

	posVelBP.Create(n)
	for posVelBP.Next() {
		velSlice := posVelVel.Slice(&posVelBP.Cursor)
		for i := range velSlice {
			velSlice[i].X, velSlice[i].Y = 1, 1
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
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
