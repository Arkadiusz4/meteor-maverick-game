package game

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"math"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: position,
		sprite:   sprite,
		rotation: 0,
	}
}

//func (p *Player) Update() {
//	speed := 5.0
//
//	if ebiten.IsKeyPressed(ebiten.KeyDown) {
//		p.position.Y += speed
//	}
//	if ebiten.IsKeyPressed(ebiten.KeyUp) {
//		p.position.Y -= speed
//	}
//	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
//		p.position.X -= speed
//	}
//	if ebiten.IsKeyPressed(ebiten.KeyRight) {
//		p.position.X += speed
//	}
//}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}

//func (p *Player) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	op.GeoM.Translate(p.position.X, p.position.Y)
//	screen.DrawImage(p.sprite, op)
//}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
