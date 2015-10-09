package main

import (
	"testing"
	"github.com/swapagarwal/gojudge"
)

func TestConfig(t *testing.T) {
	config := gojudge.GetLanguageConfig("Python")
	if config.Time != 10 {
		t.Error("failed to set time limit")
	}
	if config.Memory != 50000 {
		t.Error("failed to set memory limit")
	}
	if config.Compile != "" {
		t.Error("failed to set compile script")
	}
	if config.Run != "python FILE" {
		t.Error("failed to set run script")
	}
}
