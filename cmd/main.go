package main

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	_ "github.com/Arkadiusz4/meteor-maverick-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Game struct {
	playerPosition game.Vector
}

func (g *Game) Update() error {
	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerPosition.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerPosition.X += speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(game.PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerPosition: game.Vector{X: 100, Y: 100},
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
