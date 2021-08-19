package pipe

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math/rand"
)

const (
	screenEnd   = 800
	floor       = 295
	minDistance = 400
)

type Pipes struct {
	pipes  []*Pipe
	maxPos float64
}

type Pipe struct {
	pipeImg  *ebiten.Image
	posX     float64
	posY     float64
	collider image.Rectangle
}

func New(pipeImg *ebiten.Image, count int) *Pipes {
	var pipes = new(Pipes)
	pipes.maxPos = screenEnd
	for i := 0; i < count; i++ {
		pipes.pipes = append(pipes.pipes, NewPipe(pipeImg))
	}
	return pipes
}

func (p *Pipes) Bounds() []image.Rectangle {
	var rects []image.Rectangle
	for _, pipe := range p.pipes {
		rects = append(rects, pipe.Bounds())
	}
	return rects
}

func (p *Pipes) Draw(screen *ebiten.Image) {
	for _, pipe := range p.pipes {
		pipe.Draw(screen)
	}
}

func (p *Pipes) Update(speed float64) {
	p.maxPos -= speed
	fmt.Println(p.maxPos)
	for _, pipe := range p.pipes {
		pipe.Update(speed)

		w, _ := pipe.pipeImg.Size()
		if pipe.GetPos() <= -float64(w) {
			addDistance := screenEnd + float64(rand.Intn(3)*minDistance)
			for addDistance < p.maxPos+minDistance {
				addDistance += minDistance
			}
			pipe.SetPos(addDistance)
			p.maxPos = addDistance
		}
	}
}

func NewPipe(pipeImg *ebiten.Image) *Pipe {
	w, h := pipeImg.Size()
	return &Pipe{
		pipeImg:  pipeImg,
		posX:     screenEnd,
		posY:     floor,
		collider: image.Rect(0, 0, w/2, h),
	}
}

func (p *Pipe) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.posX, p.posY)
	screen.DrawImage(p.pipeImg, op)
}

func (p *Pipe) Update(speed float64) {
	p.Move(speed)
}

func (p *Pipe) Move(speed float64) {
	p.posX -= speed
}

func (p *Pipe) Bounds() image.Rectangle {
	rect := p.collider
	point := image.Point{
		X: int(p.posX),
		Y: int(p.posY),
	}
	rect = rect.Add(point)
	return rect
}

func (p *Pipe) SetPos(x float64) {
	p.posX = x
}

func (p *Pipe) GetPos() float64 {
	return p.posX
}

func (p *Pipe) Overlaps(rects []image.Rectangle) bool {
	rectP := p.pipeImg.Bounds()
	point := image.Point{
		X: int(p.posX),
		Y: int(p.posY),
	}
	rectP = rectP.Add(point)
	for _, rect := range rects {
		if rectP.Overlaps(rect) {
			return true
		}
	}
	return false
}
