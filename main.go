package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"image"
	_ "image/png"
	"log"
	"main/oleg_minaylow/background"
	"main/oleg_minaylow/character"
	pip "main/oleg_minaylow/pipe"
	_ "net/http/pprof"
	"os"
)

//TODO: Full game reset, freeze screen when lost (unfreeze when space is pressed),
//TODO: Make time gap between spawns

const (
	screeWidth  = 800
	screeHeight = 500
)

var (
	bg    *background.Bg
	pipes []*pip.Pipe
	char  *character.Character
)

func init() {
	bgImg, err := readImg("resources/floorBig.png")
	if err != nil {
		log.Fatal(err)
	}
	bg = background.New(bgImg)
	olegImg, err := readImg("resources/olegsamokat1.png")
	if err != nil {
		log.Fatal(err)
	}
	char = character.New(olegImg)

	pipeImg, err := readImg("resources/pinkpipe1.png")
	if err != nil {
		log.Fatal(err)
	}
	pipes = append(pipes, pip.New(pipeImg))
	pipes = append(pipes, pip.New(pipeImg))
}

func readImg(path string) (*ebiten.Image, error) {
	imgInfo, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(imgInfo)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Draw(screen *ebiten.Image) {
	bg.Draw(screen)
	for _, pipe := range pipes {
		pipe.Draw(screen)
	}
	char.Draw(screen)
}

func (g *Game) Update() error {
	bg.Update()
	char.UpdatePhysics()
	if g.isAnyKeyJustPressed() {
		go char.Action(g.keys)
	}
	char.Left()
	for _, pipe := range pipes {
		pipe.Update()
		if char.Overlaps(pipe.Bounds()) {
			char.SetPos(0, 116)
			pipe.SetPos(screeWidth)
			bg.Reset()
		}
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screeWidth, screeHeight
}

func (g *Game) isAnyKeyJustPressed() bool {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if len(g.keys) > 0 {
		return true
	}
	return false
}

func main() {
	ebiten.SetWindowSize(screeWidth, screeHeight)
	ebiten.SetWindowTitle("Check")
	ebiten.SetVsyncEnabled(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
