package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"io/ioutil"
	"path/filepath"
)

var assets embed.FS

var PlayerSprite = mustLoadImage("assets/player.png")
var LaserSprite = mustLoadImage("assets/laserBlue.png")
var MeteorSprites = mustLoadImages("assets/meteors")

func mustLoadImage(name string) *ebiten.Image {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(name)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(path string) []*ebiten.Image {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var images []*ebiten.Image
	for _, file := range files {
		if !file.IsDir() {
			imagePath := filepath.Join(path, file.Name())
			img := mustLoadImage(imagePath)
			images = append(images, img)
		}
	}

	return images
}
