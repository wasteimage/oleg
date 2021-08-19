package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	floor   = 193
	gravity = 0.32
)

type Character struct {
	olegImg *ebiten.Image
	jumping bool
	posX    float64
	posY    float64
	speedY  float64
}

func New(olegImg *ebiten.Image) *Character {
	w, _ := olegImg.Size()
	return &Character{
		olegImg: olegImg,
		posX:    float64(-w),
		posY:    floor,
	}
}

func (c *Character) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.posX, c.posY)
	screen.DrawImage(c.olegImg, op)
}
func (c *Character) Jump() {
	c.jumping = true
	defer func() { c.jumping = false }()

	c.speedY = 9.64
}
func (c *Character) Action(keys []ebiten.Key) {
	for _, key := range keys {
		if key == ebiten.KeySpace {
			if c.posY == floor {
				c.Jump()
			}
		}
	}
}

func (c *Character) UpdatePhysics() {
	c.posY -= c.speedY
	if c.posY > floor {
		c.posY = floor
	}
	c.speedY -= gravity
}

func (c *Character) Left() {
	if c.posX > 60 {
		return
	}
	c.posX += 1
}

func (c *Character) Right() {

	c.posX -= 2
}

func (c *Character) Overlaps(rects []image.Rectangle) bool {
	rectC := c.olegImg.Bounds()
	point := image.Point{
		X: int(c.posX),
		Y: int(c.posY),
	}
	rectC = rectC.Add(point)
	for _, rect := range rects {
		if rectC.Overlaps(rect) {
			return true
		}
	}
	return false
}

func (c *Character) SetPos(x, y float64) {
	c.posX, c.posY = x, y
}
