package game

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"math"
	"time"
)

type Player struct {
	game *Game

	Position Vector
	sprite   *ebiten.Image
	Rotation float64

	ShootCooldown *Timer
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		game:          game,
		Position:      position,
		sprite:        sprite,
		Rotation:      0,
		ShootCooldown: NewTimer(time.Millisecond * 500),
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Rotation += speed
	}

	p.ShootCooldown.Update()
	if p.ShootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.ShootCooldown.Reset()

		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.Position.X + halfW + math.Sin(p.Rotation)*50.0,
			p.Position.Y + halfH + math.Cos(p.Rotation)*-50.0,
		}

		bullet := NewBullet(spawnPos, p.Rotation)
		p.game.AddBullet(bullet)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.Rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.Position.X, p.Position.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return NewRect(p.Position.X, p.Position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
