package plugins

import "fmt"

type PluginVersion struct {
	Major int
	Minor int
	Micro int
}

func (pversion *PluginVersion) DisplayString() string {
	return fmt.Sprintf("%d.%d.%d", pversion.Major, pversion.Minor, pversion.Micro)
}

type PluginMetadata struct {
	Name             string
	ShortDescription string
	Author           string
	AuthorUrl        string
	Version          PluginVersion
}
