/*
    @As mentioned in PDF, this program should work like a standalone package and user 
    should be able to go get ...

    @Since im not uploading it on Git and hence i just put a placeholder :upwork/TaskBFCompiler/pkg
    @To use this in other project: go get => compiler.Compile(code)

*/
package compiler

import (
    "fmt"
    "os"
    "bufio"
    "log" 
)

// holds Raw file content to processed one
type Block struct {
    instructions []byte
    size         int
    at           int
}

/*
    @Exported Compile
    @Accepts Code(input fiele content) 
    @Executes buildFromBlock
*/
func Compile(code string) {
    // initiate codeBlock with zer values 
    var program = new(Block)

    // size :matter of processing....Depends 
    program.size = 50000
    program.instructions = make([]byte, program.size, program.size)
    program.at = 0
    buildFromBlock(program, code)
}

/*
    @Private buildFromBlock
    @Accepts Pointer to Block, code(input fiele content)
    @Returns standard output
*/
func buildFromBlock(program *Block, code string) {
    var loopStart = -1
    var loopEnd = -1
    var ignore = 0
    var skipClosingLoop = 0

    for pos, char := range code {
        if ignore == 1 {
            if char == '[' {
                skipClosingLoop += 1
            } else if char == ']' {
                if skipClosingLoop != 0 {
                    skipClosingLoop -= 1
                    continue
                }

                loopEnd = pos
                ignore = 0
                if loopStart == loopEnd {
                    loopStart = -1
                    loopEnd = -1
                    continue
                }
                loop := code[loopStart:loopEnd]
                for program.instructions[program.at] > 0 {
                    buildFromBlock(program, loop)
                }
            }
            continue
        }
        
        /*
            @At the moment, i had idea about only these operators and not sure if there are more  instructions
            @To add more: simple put it inside another `else if` block and push/pop  
        */

        if char == '+' {
            program.instructions[program.at] += 1
        } else if char == '-' {
            program.instructions[program.at] -= 1
        } else if char == '>' {
            if program.at == program.size-1 {
                program.at = 0
            } else {
                program.at += 1
            }
        } else if char == '<' {
            if program.at == 0 {
                program.at = program.size - 1
            } else {
                program.at -= 1
            }
        } else if char == '.' {
            fmt.Printf("%c", rune(program.instructions[program.at]))
        } else if char == ',' {
            fmt.Print("input: ")
            reader := bufio.NewReader(os.Stdin)
            input, _, err := reader.ReadRune()
            if err != nil{
                log.Fatalf("\x1b[%dm%s\x1b[0m", 41, err)
                return
            }

            program.instructions[program.at] = byte(input)
        } else if char == '[' {
            loopStart = pos + 1
            ignore = 1
        }
    }
}