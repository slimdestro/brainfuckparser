# Brainf*ck compiler Golang
#### github.com/slimdestro/brainfuckparser/pkg

[![N|Solid](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/60px-Go_Logo_Blue.svg.png)](https://dev.to/slimdestro)
 
## Installation

Install the package by doing go get:

```sh
go get github.com/slimdestro/brainfuckparser/pkg

compiler.Compile(syntax)  
```

## Example

```sh
package main

import (
    "github.com/slimdestro/brainfuckparser/pkg" 
    "io/ioutil"
    "errors"
    "log"
    "os" 
)
/*
    @ go run main.go bfsourcecodes/{fileName}
*/

func main() {
    if len(os.Args) > 1 {
        file, err := ioutil.ReadFile(os.Args[1])
        if err != nil{
            if errors.Is(err, os.ErrNotExist){
                log.Println("Specified file doesnt exist!")
                return
            }else{
                panic(err)
            }
        }

        code := string(file) 
        compiler.Compile(code)
    } else {
        log.Fatal("a BrainFcuk source file is required")
    }
}
 
```


## Author

[slimdestro(Mukul Mishra)](https://linktr.ee/slimdestro)
