package tests

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	"math"
	"testing"
)

func TestNormalize(t *testing.T) {
	v1 := game.Vector{X: 3, Y: 4}
	normalizedV1 := v1.Normalize()
	expectedV1 := game.Vector{X: 0.6, Y: 0.8}
	if !vectorEqual(normalizedV1, expectedV1) {
		t.Errorf("Expected %v, but got %v", expectedV1, normalizedV1)
	}

	v2 := game.Vector{X: -3, Y: -4}
	normalizedV2 := v2.Normalize()
	expectedV2 := game.Vector{X: -0.6, Y: -0.8}
	if !vectorEqual(normalizedV2, expectedV2) {
		t.Errorf("Expected %v, but got %v", normalizedV2, normalizedV1)
	}
}

func vectorEqual(v1, v2 game.Vector) bool {
	epsilon := 0.000001
	return math.Abs(v1.X-v2.X) < epsilon && math.Abs(v1.Y-v2.Y) < epsilon
}
