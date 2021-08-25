package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"image"
	_ "image/png"
	"log"
	_ "net/http/pprof"
	"oleg/game"
	"os"
)

//TODO: Add sounds

var greenHillImg, nightCityImg, egyptImg, olegImg, olegEgyptImg, pipeGreenHillImg, pipeNightCityImg, pipeEgyptImg, loseImg, icon *ebiten.Image
var music *audio.Player
var scorePath = "resources/best_score"

const (
	baseSpeed  = 4.
	sampleRate = 48000
)

func init() {
	var err error
	greenHillImg, err = readImg("resources/floorBig.png")
	if err != nil {
		log.Fatal(err)
	}
	nightCityImg, err = readImg("resources/floor2LocNightCity.png")
	if err != nil {
		log.Fatal(err)
	}
	egyptImg, err = readImg("resources/floor3LocEgypt.png")
	if err != nil {
		log.Fatal(err)
	}
	olegImg, err = readImg("resources/olegsamokat1.png")
	if err != nil {
		log.Fatal(err)
	}
	olegEgyptImg, err = readImg("resources/olegsamokat3LocEgypt.png")
	if err != nil {
		log.Fatal(err)
	}
	pipeGreenHillImg, err = readImg("resources/pinkpipe1.png")
	if err != nil {
		log.Fatal(err)
	}
	pipeNightCityImg, err = readImg("resources/pinkpipe2LocNightCity.png")
	if err != nil {
		log.Fatal(err)
	}
	pipeEgyptImg, err = readImg("resources/pinkpipe3LocEgypt.png")
	if err != nil {
		log.Fatal(err)
	}
	loseImg, err = readImg("resources/press_any_key.png")
	if err != nil {
		log.Fatal(err)
	}
	icon, err = readImg("resources/icon.png")
	if err != nil {
		log.Fatal(err)
	}
	music, err = readMp3("resources/audio/oleg_minaylow_game_melody.mp3")
	if err != nil {
		log.Fatal(err)
	}
}

func readImg(path string) (*ebiten.Image, error) {
	imgInfo, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imgInfo.Close()
	img, _, err := image.Decode(imgInfo)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
func readMp3(path string) (*audio.Player, error) {
	audioCtx := audio.NewContext(sampleRate)
	mp3Info, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	msc, err := mp3.DecodeWithSampleRate(sampleRate, mp3Info)
	if err != nil {
		return nil, err
	}
	np, err := audio.NewPlayer(audioCtx, msc)
	return np, err
}

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Oleg Minaylow Travel")
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowIcon([]image.Image{
		icon,
	})
	if err := ebiten.RunGame(game.New(
		greenHillImg, nightCityImg, egyptImg, olegImg, olegEgyptImg, pipeGreenHillImg, pipeNightCityImg, pipeEgyptImg, loseImg,
		scorePath,
		baseSpeed,
		music,
	)); err != nil {
		log.Fatal(err)
	}
	music.Close()
}
