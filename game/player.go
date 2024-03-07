package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

var PlayerSprite = mustLoadImage("assets/player.png")

func mustLoadImage(name string) *ebiten.Image {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(name)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
