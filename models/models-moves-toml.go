package models

type MovesEditor struct {
	Move []Move `toml:"move"`
}
type Descriptions struct {
	Description string `toml:"description"`
}
type Effects struct {
	EffectText string `toml:"effect_text"`
}
type Move struct {
	Name         string         `toml:"name"`
	Power        int            `toml:"Power"`
	Pp           int            `toml:"pp"`
	Accuracy     int            `toml:"accuracy"`
	Type         string         `toml:"type"`
	KindOfMove   string         `toml:"kind_of_move"`
	Descriptions []Descriptions `toml:"descriptions"`
	Effects      []Effects      `toml:"effects"`
}
