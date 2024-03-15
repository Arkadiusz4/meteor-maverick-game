package game

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Bullet struct {
	Position Vector
	Rotation float64
	Sprite   *ebiten.Image
}

func NewBullet(position Vector, rotation float64) *Bullet {
	sprite := assets.LaserSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfH

	return &Bullet{
		Position: position,
		Rotation: rotation,
		Sprite:   sprite,
	}
}

func (b *Bullet) Update() {
	bulletSpeed := 350.0 / float64(ebiten.TPS())

	b.Position.X += math.Sin(b.Rotation) * bulletSpeed
	b.Position.Y += math.Cos(b.Rotation) * -bulletSpeed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	bounds := b.Sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.Rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.Position.X, b.Position.Y)

	screen.DrawImage(b.Sprite, op)
}

func (b *Bullet) Collider() Rect {
	bounds := b.Sprite.Bounds()

	return NewRect(b.Position.X, b.Position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
