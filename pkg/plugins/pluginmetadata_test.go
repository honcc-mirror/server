package plugins

import (
	"testing"
)

func TestPluginVersionDisplayString(t *testing.T) {
	pluginVersion := PluginVersion{Major: 3, Minor: 2, Micro: 1}
	displayString := pluginVersion.DisplayString()
	expectedDisplayString := "3.2.1"

	if displayString != expectedDisplayString {
		t.Fatalf(`Did not receive expected version display string.
			Received: %s
			Expected: %s`, displayString, expectedDisplayString)
	}
}
