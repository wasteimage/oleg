package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	_ "net/http/pprof"
	"oleg/game"
	"os"
)

//TODO: Change LVLs, add music & sounds

var bgImg, olegImg, pipeImg, loseImg *ebiten.Image
var scorePath = "resources/best_score"

const baseSpeed = 4.

func init() {
	var err error
	bgImg, err = readImg("resources/floor2LocNightCity.png")
	if err != nil {
		log.Fatal(err)
	}
	olegImg, err = readImg("resources/olegsamokat1.png")
	if err != nil {
		log.Fatal(err)
	}
	pipeImg, err = readImg("resources/pinkpipe2LocNightCity.png")
	if err != nil {
		log.Fatal(err)
	}
	loseImg, err = readImg("resources/press_any_key.png")
	if err != nil {
		log.Fatal(err)
	}
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

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("OLEG MINAYLOW TRAVEL")
	ebiten.SetVsyncEnabled(true)
	if err := ebiten.RunGame(game.New(olegImg, pipeImg, bgImg, loseImg, scorePath, baseSpeed)); err != nil {
		log.Fatal(err)
	}
}
