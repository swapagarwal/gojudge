package gojudge

import (
    "io/ioutil"
    "github.com/antonholmquist/jason"
)

type LanguageConfig struct {
    Time int64
    Memory int64
    Compile string
    Run string
}

func GetLanguageConfig(lang string) LanguageConfig {
    raw, _ := ioutil.ReadFile("./config.json")
    decoded, _ := jason.NewObjectFromBytes([]byte(raw))
    time, _ := decoded.GetInt64("Languages", lang, "Time")
    memory, _ := decoded.GetInt64("Languages", lang, "Memory")
    compile, _ := decoded.GetString("Languages", lang, "Compile")
    run, _ := decoded.GetString("Languages", lang, "Run")
    return LanguageConfig{
        Time: time,
        Memory: memory,
        Compile: compile,
        Run: run,
    }
}
