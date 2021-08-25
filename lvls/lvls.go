package lvls

type Lvl int

const (
	LvlGreenHill Lvl = iota
	LvlNightCity
	LvlEgypt
)

func CurrentLvl(gameTime float64) Lvl {
	switch {
	case gameTime > 40:
		return LvlEgypt
	case gameTime > 20:
		return LvlNightCity
	default:
		return LvlGreenHill
	}
}
