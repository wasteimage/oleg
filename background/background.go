package background

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bg struct {
	bgImg *ebiten.Image
	posX  float64
}

func New(bgImg *ebiten.Image) *Bg {
	return &Bg{bgImg: bgImg}
}

func (b *Bg) Draw(screen *ebiten.Image) {
	w, _ := b.bgImg.Size()
	op1 := new(ebiten.DrawImageOptions)
	op1.GeoM.Translate(b.posX, 0)
	screen.DrawImage(b.bgImg, op1)

	op2 := new(ebiten.DrawImageOptions)
	op2.GeoM.Translate(b.posX+float64(w), 0)
	screen.DrawImage(b.bgImg, op2)

}

func (b *Bg) Update() {
	b.Move()
	w, _ := b.bgImg.Size()
	if b.posX <= -float64(w) {
		b.posX = 0
	}
}

func (b *Bg) Move() {
	b.posX -= 4
}
