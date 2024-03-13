package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
	"path/filepath"
)

var assets embed.FS

var PlayerSprite = mustLoadImage("assets/player.png")
var LaserSprite = mustLoadImage("assets/laserBlue.png")
var MeteorSprites = mustLoadImages("assets/meteors")
var ScoreFont = mustLoadFont("assets/font.ttf")

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

func mustLoadFont(filePath string) font.Face {
	fontData, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	parseFont, err := opentype.Parse(fontData)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(parseFont, &opentype.FaceOptions{
		Size: 48,
		DPI:  72,
	})
	if err != nil {
		panic(err)
	}

	return face
}
