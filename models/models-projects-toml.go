package models

type Projects struct {
	Project []Project `toml:"project"`
}
type Project struct {
	Name            string `toml:"name"`
	FolderLocation  string `toml:"folderLocation"`
	VersionOfEngine string `toml:"versionOfEngine"`
	CreatedDateTime string `toml:"createdDateTime"`
	LastUsed        string `toml:"lastedUsed"`
	ID              string `toml:"id"`
}
