package background

import (
	"github.com/hajimehoshi/ebiten/v2"
	"oleg/lvls"
)

type Bg struct {
	bgImgs map[lvls.Lvl]*ebiten.Image
	posX   float64
}

func New(greenHillImg, nightCityImg, egyptImg *ebiten.Image) *Bg {
	return &Bg{bgImgs: map[lvls.Lvl]*ebiten.Image{
		lvls.LvlGreenHill: greenHillImg,
		lvls.LvlNightCity: nightCityImg,
		lvls.LvlEgypt:     egyptImg,
	}}
}

func (b *Bg) Draw(screen *ebiten.Image, lvl lvls.Lvl) {
	var bgImg = b.bgImgs[lvl]
	w, _ := bgImg.Size()
	op1 := new(ebiten.DrawImageOptions)
	op1.GeoM.Translate(b.posX, 0)
	screen.DrawImage(bgImg, op1)

	op2 := new(ebiten.DrawImageOptions)
	op2.GeoM.Translate(b.posX+float64(w), 0)
	screen.DrawImage(bgImg, op2)

}

func (b *Bg) Update(speed float64, lvl lvls.Lvl) {
	var bgImg = b.bgImgs[lvl]
	b.Move(speed)
	w, _ := bgImg.Size()
	if b.posX <= -float64(w) {
		b.posX = 0
	}
}

func (b *Bg) Move(speed float64) {
	b.posX -= speed * 0.5
}
