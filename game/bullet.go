package game

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Bullet struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewBullet(position Vector, rotation float64) *Bullet {
	sprite := assets.LaserSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfH

	return &Bullet{
		position: position,
		rotation: rotation,
		sprite:   sprite,
	}
}

func (b *Bullet) Update() {
	bulletSpeed := 350.0 / float64(ebiten.TPS())

	b.position.X += math.Sin(b.rotation) * bulletSpeed
	b.position.Y += math.Cos(b.rotation) * -bulletSpeed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	bounds := b.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.sprite, op)
}

func (b *Bullet) Collider() Rect {
	bounds := b.sprite.Bounds()

	return NewRect(b.position.X, b.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
