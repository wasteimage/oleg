package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"oleg/background"
	"oleg/character"
	pip "oleg/pipe"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 500
)

var speed = 0

type Game struct {
	keys      []ebiten.Key
	state     State
	resetGame func(*Game)

	Lose    bool
	loseImg *ebiten.Image
}

type State struct {
	char  *character.Character
	pipes []*pip.Pipe
	bg    *background.Bg
}

func New(olegImg, pipeImg, bgImg, loseImg *ebiten.Image) (g *Game) {
	g = &Game{
		resetGame: func(game *Game) {
			var (
				bg    *background.Bg
				char  *character.Character
				pipes []*pip.Pipe
			)
			bg = background.New(bgImg)
			char = character.New(olegImg)
			pipes = append(pipes, pip.New(pipeImg))
			pipes = append(pipes, pip.New(pipeImg))
			game.state = State{
				char:  char,
				pipes: pipes,
				bg:    bg,
			}
		},
		loseImg: loseImg,
	}
	g.resetGame(g)
	return g
}

func (g *Game) ResetGame(game *Game) {
	g.resetGame(game)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.bg.Draw(screen)
	for _, pipe := range g.state.pipes {
		pipe.Draw(screen)
	}
	g.state.char.Draw(screen)
	if g.Lose {
		g.LoseScreen(screen)
	}
}

func (g *Game) Update() error {
	if g.Lose {
		g.state.bg.StopTimer()
		g.state.bg.CheckMaxScore()
		if g.isAnyKeyJustPressed() {
			g.Lose = false
			g.ResetGame(g)
		}
		return nil
	}
	g.state.bg.Update()
	g.state.char.UpdatePhysics()
	if g.isAnyKeyJustPressed() {
		go g.state.char.Action(g.keys)
	}
	g.state.char.Left()
	for _, pipe := range g.state.pipes {
		pipe.Update()
		if g.state.char.Overlaps(pipe.Bounds()) {
			g.Lose = true
		}
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
