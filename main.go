package main

import (
    "github.com/slimdestro/brainfuckparser/pkg" 
    "io/ioutil"
    "errors"
    "log"
    "os" 
)
/*
    @please run this file : go run main.go bfsourcecodes/{fileName}
*/

func main() {
    if len(os.Args) > 1 {
        file, err := ioutil.ReadFile(os.Args[1])
        if err != nil{
            if errors.Is(err, os.ErrNotExist){
                log.Printf("\x1b[%dm%s\x1b[0m", 41, "Specified file doesnt exist!")
                return
            }else{
                panic(err)
            }
        }

        code := string(file) 
        compiler.Compile(code)
    } else {
        log.Fatalf("\x1b[%dm%s\x1b[0m", 41,"Unfortunately a BrainFcuk source file is required :)")
    }
}

