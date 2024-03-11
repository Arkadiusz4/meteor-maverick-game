package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player          *Player
	meteorSpawnTime *Timer
	meteors         []*Meteor
}

func NewGame() *Game {
	g := &Game{}

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

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
