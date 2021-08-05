package background

import (
	"context"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"strconv"
	"time"
)

type Bg struct {
	bgImg  *ebiten.Image
	posX   float64
	time   int
	cancel context.CancelFunc
}

func New(bgImg *ebiten.Image) *Bg {
	b := &Bg{bgImg: bgImg}
	b.StartTimer()
	return b
}

func (b *Bg) Draw(screen *ebiten.Image) {
	w, _ := b.bgImg.Size()
	op1 := new(ebiten.DrawImageOptions)
	op1.GeoM.Translate(b.posX, 0)
	screen.DrawImage(b.bgImg, op1)

	op2 := new(ebiten.DrawImageOptions)
	op2.GeoM.Translate(b.posX+float64(w), 0)
	screen.DrawImage(b.bgImg, op2)
	ebitenutil.DebugPrint(screen, strconv.Itoa(b.time))
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

func (b *Bg) StartTimer() {
	var ctx, cancel = context.WithCancel(context.Background())
	b.cancel = cancel
	ticker := time.NewTicker(time.Second)
	go func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				b.time++
			case <-ctx.Done():
				ticker.Stop()
				b.time = 0
				return
			}
		}
	}(ctx)
}

func (b *Bg) Reset() {
	b.cancel()
	b.StartTimer()
}
