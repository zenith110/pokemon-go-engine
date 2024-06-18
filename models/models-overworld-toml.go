package models

type OverworldsHolder struct {
	Overworlds []Overworld `toml:"overworld"`
}
type OverworldDirectionFrame struct {
	Direction   string `toml:"direction"`
	FrameNumber string `toml:"frameNumber"`
	Path        string `toml:"path"`
	Sprite      string `toml:"sprite"`
}
type Swimming struct {
	OverworldDirectionFrames []OverworldDirectionFrame `toml:"OverworldDirectionFrames"`
}

type Surfing struct {
	OverworldDirectionFrames []OverworldDirectionFrame `toml:"OverworldDirectionFrames"`
}

type Walking struct {
	OverworldDirectionFrames []OverworldDirectionFrame `toml:"OverworldDirectionFrames"`
}
type Running struct {
	OverworldDirectionFrames []OverworldDirectionFrame `toml:"OverworldDirectionFrames"`
}
type Overworld struct {
	Name     string   `toml:"name"`
	ID       string   `toml:"id"`
	IsPlayer bool     `toml:"isPlayer"`
	Swimming Swimming `toml:"swimming"`
	Walking  Walking  `toml:"walking"`
	Running  Running  `toml:"running"`
	Surfing  Surfing  `toml:"surfing"`
}
