package game

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
