package game

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
)

type Meteor struct {
	position Vector
	movement Vector
	rotation float64
	sprite   *ebiten.Image
}

//var meteorSprites = assets.MeteorSprites

func NewMeteor() *Meteor {
	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	r := ScreenWidth / 2.0

	angle := rand.Float64() * 2 * math.Pi

	pos := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	velocity := 0.25 + rand.Float64()*1.5

	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	//sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	if len(assets.MeteorSprites) == 0 {
		// Handle the case where there are no sprites available
		panic("No meteor sprites available")
	}

	// Generate a random index within the range of available sprites
	spriteIndex := rand.Intn(len(assets.MeteorSprites))
	sprite := assets.MeteorSprites[spriteIndex]

	return &Meteor{
		position: pos,
		movement: movement,
		sprite:   sprite,
	}
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += -0.02 + rand.Float64()*0.04
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}
