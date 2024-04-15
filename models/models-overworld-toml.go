package models

type OverworldsHolder struct {
	Overworlds []Overworld `toml:"overworld"`
}
type Up struct {
	ID        int    `toml:"id"`
	FrameName string `toml:"frameName"`
}
type Down struct {
	ID        int    `toml:"id"`
	FrameName string `toml:"frameName"`
}
type Right struct {
	ID        int    `toml:"id"`
	FrameName string `toml:"frameName"`
}
type Left struct {
	ID        int    `toml:"id"`
	FrameName string `toml:"frameName"`
}
type Swimming struct {
	Up    []Up    `toml:"up"`
	Down  []Down  `toml:"down"`
	Right []Right `toml:"right"`
	Left  []Left  `toml:"left"`
}
type Walking struct {
	Up    []Up    `toml:"up"`
	Down  []Down  `toml:"down"`
	Left  []Left  `toml:"left"`
	Right []Right `toml:"right"`
}
type Running struct {
	Up    []Up    `toml:"up"`
	Down  []Down  `toml:"down"`
	Left  []Left  `toml:"left"`
	Right []Right `toml:"right"`
}
type Overworld struct {
	Name     string   `toml:"name"`
	ID       string   `toml:"id"`
	IsPlayer bool     `toml:"isPlayer"`
	Swimming Swimming `toml:"swimming"`
	Walking  Walking  `toml:"walking"`
	Running  Running  `toml:"running"`
}
