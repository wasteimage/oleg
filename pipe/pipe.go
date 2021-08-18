package pipe

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math/rand"
	"time"
)

const (
	screenEnd = 800
	floor     = 295
)

type Pipe struct {
	pipeImg   *ebiten.Image
	posX      float64
	posY      float64
	created   time.Time
	timeStart time.Time
	collider  image.Rectangle
}

func New(pipeImg *ebiten.Image) *Pipe {
	w, h := pipeImg.Size()
	return &Pipe{
		pipeImg:   pipeImg,
		posX:      screenEnd,
		posY:      floor,
		created:   time.Now().Add(time.Duration(1+rand.Intn(3)) * time.Second),
		timeStart: time.Now(),
		collider:  image.Rect(0, 0, w/2, h),
	}
}

func (p *Pipe) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.posX, p.posY)
	screen.DrawImage(p.pipeImg, op)
}

func (p *Pipe) Update(speed float64) {
	if time.Now().Before(p.created) {
		return
	}
	p.Move(speed)
	w, _ := p.pipeImg.Size()
	if p.posX <= -(0 + float64(w)) {
		p.posX = screenEnd
		p.created = time.Now().Add(time.Duration(1+rand.Intn(3)) * time.Second)
	}
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
