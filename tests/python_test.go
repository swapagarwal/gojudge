package main

import (
    "io/ioutil"
    "os"
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
    if config.Compile != ":" {
        t.Error("failed to set compile script")
    }
    if config.Run != "python FILE" {
        t.Error("failed to set run script")
    }
    if config.FilePlaceholder != "FILE" {
        t.Error("failed to set file placeholder")
    }
}

func TestMemory(t *testing.T) {
    config := gojudge.GetLanguageConfig("Python")
    gojudge.Run(config.Time, config.Memory, config.Compile, config.Run, config.FilePlaceholder, "scripts/memory.py")
    out, err := ioutil.ReadFile("out.txt")
    if err != nil {
        t.Error("failed to open output file")
    }
    if string(out) == "Done!\n" {
        t.Error("failed to limit memory")
    }
}

func TestSum(t *testing.T) {
    config := gojudge.GetLanguageConfig("Python")
    in, err := os.Create("in.txt")
    if err != nil {
        t.Error("failed to create input file")
    }
    _, err = in.WriteString("2\n3\n")
    if err != nil {
        t.Error("failed to write input file")
    }
    in.Sync()
    in.Close()
    gojudge.Run(config.Time, config.Memory, config.Compile, config.Run, config.FilePlaceholder, "scripts/sum.py")
    out, err := ioutil.ReadFile("out.txt")
    if err != nil {
        t.Error("failed to open output file")
    }
    if string(out) != "5\n" {
        t.Error("failed to compute the sum")
    }
}

func TestTime(t *testing.T) {
    config := gojudge.GetLanguageConfig("Python")
    gojudge.Run(config.Time, config.Memory, config.Compile, config.Run, config.FilePlaceholder, "scripts/time.py")
    out, err := ioutil.ReadFile("out.txt")
    if err != nil {
        t.Error("failed to open output file")
    }
    if string(out) == "Done!\n" {
        t.Error("failed to limit time")
    }
}
