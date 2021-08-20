package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"oleg/background"
	"oleg/character"
	"oleg/lvls"
	pip "oleg/pipe"
	"oleg/score"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 500
)

type Game struct {
	keys      []ebiten.Key
	state     State
	resetGame func(*Game)

	Lose    bool
	loseImg *ebiten.Image

	baseSpeed float64

	lvl lvls.Lvl
}

type State struct {
	char  *character.Character
	pipes *pip.Pipes
	bg    *background.Bg
	score *score.Score
	speed float64
}

func New(
	greenHillImg, nightCityImg, egyptImg, olegImg, olegEgyptImg, pipeGreenHillImg, pipeNightCityImg, pipeEgyptImg, loseImg *ebiten.Image,
	scorePath string,
	baseSpeed float64,
) (g *Game) {
	g = &Game{
		resetGame: func(game *Game) {
			var (
				bg    *background.Bg
				char  *character.Character
				pipes *pip.Pipes
				scr   *score.Score
			)
			bg = background.New(greenHillImg, nightCityImg, egyptImg)
			char = character.New(olegImg, olegEgyptImg)
			scr = score.New(scorePath)
			pipes = pip.New(pipeGreenHillImg, pipeNightCityImg, pipeEgyptImg, 2)
			game.state = State{
				char:  char,
				pipes: pipes,
				bg:    bg,
				score: scr,
			}
		},
		loseImg:   loseImg,
		baseSpeed: baseSpeed,
		lvl:       lvls.LvlGreenHill,
	}
	g.resetGame(g)
	return g
}

func (g *Game) ResetGame(game *Game) {
	g.resetGame(game)
	if g.lvl != lvls.LvlGreenHill {
		g.lvl = lvls.LvlGreenHill
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.bg.Draw(screen, g.lvl)
	g.state.pipes.Draw(screen, g.lvl)
	g.state.char.Draw(screen, g.lvl)
	if g.Lose {
		g.LoseScreen(screen)
	}
	g.state.score.Draw(screen)
}

func (g *Game) Update() error {
	g.state.speed = g.baseSpeed + g.state.score.GameTime()/10
	if g.Lose {
		g.state.score.StopTimer()
		g.state.score.UpdateMaxScore()
		if g.isAnyKeyJustPressed() {
			g.Lose = false
			g.ResetGame(g)
		}
		return nil
	}
	if g.lvl != lvls.LvlNightCity && g.state.score.GameTime() > 20 {
		g.lvl = lvls.LvlNightCity
	}
	if g.lvl != lvls.LvlEgypt && g.state.score.GameTime() > 40 {
		g.lvl = lvls.LvlEgypt
	}
	g.state.bg.Update(g.state.speed, g.lvl)
	g.state.char.UpdatePhysics()
	if g.isAnyKeyJustPressed() {
		go g.state.char.Action(g.keys)
	}
	g.state.char.Left()
	g.state.pipes.Update(g.state.speed, g.lvl)
	if g.state.char.Overlaps(g.state.pipes.Bounds(), g.lvl) {
		g.Lose = true
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) isAnyKeyJustPressed() bool {
	g.keys = inpututil.PressedKeys()
	if len(g.keys) > 0 {
		return true
	}
	return false
}

func (g *Game) LoseScreen(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 15)
	screen.DrawImage(g.loseImg, op)
}
