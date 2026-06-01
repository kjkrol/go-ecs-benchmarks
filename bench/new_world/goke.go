package newworld

import (
	"testing"

	"github.com/kjkrol/goke"
	"github.com/stretchr/testify/assert"
)

func runGOKe(b *testing.B, _ int) {
	var ecs *goke.ECS
	for b.Loop() {
		ecs = goke.New()
	}
	assert.NotNil(b, ecs)
}
