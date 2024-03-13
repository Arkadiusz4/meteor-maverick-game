package game

import (
	"fmt"
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player          *Player
	meteorSpawnTime *Timer
	meteors         []*Meteor
	bullets         []*Bullet

	score int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTime: NewTimer(1 * time.Second),
	}

	g.player = NewPlayer(g)

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTime.Update()
	if g.meteorSpawnTime.IsReady() {
		g.meteorSpawnTime.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
			}
		}
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.score = 0
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
