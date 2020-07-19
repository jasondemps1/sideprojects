package main

import (
	"image"
	"bytes"
	_ "image/png"
	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}


func main() {
	img, _, err := image.Decode()//bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("title")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
