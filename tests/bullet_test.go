package tests

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	"math"
	"testing"
)

func TestNewBullet(t *testing.T) {
	position := game.Vector{X: 100, Y: 100}
	rotation := math.Pi / 4
	bullet := game.NewBullet(position, rotation)

	sprite := assets.LaserSprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfH

	if bullet.Position != position {
		t.Errorf("Expected position %v, but got %v", position, bullet.Position)
	}
	if bullet.Rotation != rotation {
		t.Errorf("Expected rotation %f, but got %f", rotation, bullet.Rotation)
	}
}

func TestBulletUpdate(t *testing.T) {
	bullet := &game.Bullet{
		Position: game.Vector{X: 100, Y: 100},
		Rotation: math.Pi / 4,
	}
	bullet.Update()
}
