package tests

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	"testing"
)

func TestRectMaxX(t *testing.T) {
	rect := game.NewRect(1, 2, 3, 4)
	expectedMaxX := 4.0
	maxX := rect.MaxX()
	if maxX != expectedMaxX {
		t.Errorf("Expected MaxX to be %f, but got %f", expectedMaxX, maxX)
	}
}

func TestRectMaxY(t *testing.T) {
	rect := game.NewRect(1, 2, 3, 4)
	expectedMaxY := 6.0
	maxY := rect.MaxY()
	if maxY != expectedMaxY {
		t.Errorf("Expected MaxY to be %f, but got %f", expectedMaxY, maxY)
	}
}

func TestRectIntersects(t *testing.T) {
	rect1 := game.NewRect(1, 2, 3, 4)
	rect2 := game.NewRect(3, 4, 3, 4)
	rect3 := game.NewRect(10, 10, 1, 1)

	if !rect1.Intersects(rect2) {
		t.Errorf("Expected rect1 and rect2 to intersect, but they do not")
	}

	if rect1.Intersects(rect3) {
		t.Errorf("Expected rect1 and rect3 not to intersect, but they do")
	}
}
