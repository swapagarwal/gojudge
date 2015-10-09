package gojudge

import (
    "io/ioutil"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "github.com/antonholmquist/jason"
)

type LanguageConfig struct {
    Time int64
    Memory int64
    Compile string
    Run string
    FilePlaceholder string
}

func GetLanguageConfig(lang string) LanguageConfig {
    raw, _ := ioutil.ReadFile("./config.json")
    decoded, _ := jason.NewObjectFromBytes([]byte(raw))
    time, _ := decoded.GetInt64("Languages", lang, "Time")
    memory, _ := decoded.GetInt64("Languages", lang, "Memory")
    compile, _ := decoded.GetString("Languages", lang, "Compile")
    run, _ := decoded.GetString("Languages", lang, "Run")
    file, _ := decoded.GetString("FileName")
    return LanguageConfig{
        Time: time,
        Memory: memory,
        Compile: compile,
        Run: run,
        FilePlaceholder: file,
    }
}

func Run(time int64, memory int64, compile string, run string, filePlaceholder string, file string) {
    app := "sh"
    arg := "-c"
    arg1 := "ulimit -St " + strconv.FormatInt(time, 10)
    arg2 := "ulimit -Sp 500"
    arg3 := "ulimit -Sn 10"
    arg4 := "ulimit -Sv " + strconv.FormatInt(memory, 10)
    arg5 := strings.Replace(compile, filePlaceholder, file, 1)
    arg6 := strings.Replace(run, filePlaceholder, file, 1)
    args := strings.Join([]string{arg1, arg2, arg3, arg4, arg5, arg6}, "; ")

    cmd := exec.Command(app, arg, args)
    cmd.Stdin, _ = os.Open("in.txt")
    cmd.Stdout, _ = os.Create("out.txt")
    cmd.Stderr, _ = os.Create("err.txt")
    err := cmd.Run()

    if err != nil {
        println(err.Error())
        return
    }
}
