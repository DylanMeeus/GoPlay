package main


// main package for alewig


import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "go/types"
    "../DataStructures/Collections/"
)

// set up section constants
const codesection string = ".code"
const datasection string = ".data"

type TokenFunction func(*CodeStack, FunctionArgs)

type CodeStack struct{
    collections.Stack
}

type FunctionArgs []interface{}

var TokenFunctions = map[Token] TokenFunction{
    Token{representation:"exit"}: exit,
    Token{representation:"push"}: push,
    Token{representation:"print"}: print,
    Token{representation:"pop"}: pop,
}

func exit(codeStack *CodeStack, args FunctionArgs){
    os.Exit(1)
}

func push(codeStack *CodeStack, args FunctionArgs){
    for i := 0; i < len(args); i++ {
        codeStack.Push(args[i])
    }
}


func pop(codeStack *CodeStack, args FunctionArgs){
    codeStack.Pop()
}

/*
    Print the top of the stack
 */
func print(codeStack *CodeStack, args FunctionArgs){
    top := codeStack.Peek()
    fmt.Print(top)
}

func main(){
    // todo: implement this with flags, optional linker in the future to native Go
    if len(os.Args) < 2{
        panic("No source file specified!")
    }
    filename := os.Args[1]
    byteContent, err := ioutil.ReadFile(filename)
    if err != nil{
        panic("Could not load file!")
    }

    parse(string(byteContent))
}

type Token struct{
    representation string       // string representation
}

type TokenLine struct{
    tokens []Token
    variableDeclaration bool               // Does this line contain variables, or code?
}


type Variable struct{
    datatype types.BasicKind
    value interface{}
}


/*
    Parse the incoming file, create an internal representations of the code.
    Since it is evaluated line by line, and we have goto statements, line numbers are important :-)
 */
func parse(source string){
    lines := strings.Split(source,"\n")
    lines = preParseFormat(lines)
    tokenLines := tokenize(lines)

    variables := make(map[string]Variable,0)
    stack := CodeStack{} // stack containing the values!
    stack.Push(1)
    stack.Push(2)

    for i := 0; i < len(tokenLines); i++ {
        tokenLine := tokenLines[i]
        if tokenLine.variableDeclaration {
            // create the variable map
            // each line should contain 2 tokens, the variable name + datatype
            name := tokenLine.tokens[0]
            dtype:= tokenLine.tokens[1]
            basicKind, defaultValue := typeFromString(dtype.representation)
            variable := Variable{datatype:basicKind, value:defaultValue}
            variables[name.representation] = variable
        } else {
            // parse the code
            token := tokenLine.tokens[0]
            function := TokenFunctions[token]
            // maybe we have some args!
            args := make([]interface{},0)
            for i:= 1; i < len(tokenLine.tokens); i++{
                args = append(args, tokenLine.tokens[i])
            }
            function(&stack, args)
        }
    }
}


/*
    Returns the type from a string, along with it's default value!
 */
func typeFromString(typestring string) (types.BasicKind, interface{}) { // only allow basic types?
    switch typestring{ // todo: maybe the spaces should be trimmed earlier!
        case "string":
            return types.String, ""
        default:
            panic("Type not supported")
    }
}

func tokenize(source []string) []TokenLine{
    parsingData := true    // assume we have a data section
    tokenLines := make([]TokenLine, 0)
    for i := 0; i < len(source); i++ {
        line := source[i]
        // parse the actual line if it was not a 'switch context statement' (.code / .data)
        if line == codesection {
            parsingData = false
            continue
        }
        if line == datasection {
            parsingData = true
            continue
        }
        if parsingData {
            // split the line based on ":"
            parts := strings.Split(line, ":")
            if len(parts) != 2{
                panic("Could not parse data section!")
            }
            variableName := Token{representation:strings.TrimSpace(parts[0])}
            variableType := Token{representation:strings.TrimSpace(parts[1])}
            tokenLine := TokenLine{tokens:[]Token{variableName, variableType},variableDeclaration:true}
            tokenLines = append(tokenLines, tokenLine)
        } else {
            parts := strings.Split(line," ")
            tokens := make([]Token,0)
            for i := 0; i < len(parts); i++ {
                part := parts[i]
                token := Token{representation:strings.TrimSpace(part)}
                tokens = append(tokens, token)
            }
            tokenLines = append(tokenLines, TokenLine{tokens: tokens, variableDeclaration:false})
        }
    }
    return tokenLines
}

/*
    Format the code before parsing to remove some unneccessary elements
    e.g: No empty lines, comments, ..
 */
func preParseFormat(source []string) []string{ // keep this side-effect free, take the perf hit.
    deleteComments(&source)
    deleteEmptyLines(&source)
    deleteIndentation(&source) // we're not python! :-)
    return source
}

func deleteComments(source *[]string){
    for i := 0; i < len(*source); i++{
        line := (*source)[i]
        index := strings.Index(line, "//")
        if index != -1 {
            (*source)[i] = line[0:index]
        }
    }
}

func deleteEmptyLines(source *[]string) {
    contentLines := make([]string, 0)
    for i := 0; i < len(*source); i++{
        line := (*source)[i]
        if len(strings.TrimSpace(line)) != 0{
            contentLines = append(contentLines, line)
        }
    }
    *source = contentLines
}

func deleteIndentation(source *[]string){
    var spacerune rune
    spacerune = 32

    for i := 0; i < len(*source); i++ {
        line := (*source)[i]
        line = strings.Map(func(r rune) rune{
            if r == spacerune{
                return -1
            }
            return r
        },line)
    }
}

func printSource(source []string){
    for i := 0; i < len(source); i++ {
        fmt.Println(source[i])
    }
}