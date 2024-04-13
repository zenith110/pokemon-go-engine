package models

type Projects struct {
	Project struct {
		Name            string `toml:"name"`
		FolderLocation  string `toml:"folderLocation"`
		VersionOfEngine string `toml:"versionOfEngine"`
		CreatedDateTime string `toml:"createdDateTime"`
	} `toml:"project"`
}
