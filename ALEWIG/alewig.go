package main


// main package for alewig


import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "go/types"
    "../DataStructures/Collections/"
    "strconv"
)

// set up section constants
const codesection string = ".code"
const datasection string = ".data"

var variables = make(map[string]Variable,0)
var jumpTable = make(map[string]uint,0)     // jump table indexed by string

type TokenFunction func(*CodeStack, *uint, FunctionArgs)

type CodeStack struct{
    collections.Stack
}

type FunctionArgs []interface{}

var TokenFunctions = map[Token] TokenFunction{
    Token{representation:"assign"}: assign,
    Token{representation:"exit"}: exit,
    Token{representation:"print"}: print,

    // stack operations
    Token{representation:"push"}: push,
    Token{representation:"pop"}: pop,
    Token{representation:"decrement"} : decrement, // -- on top of stack
    Token{representation:"increment"} : increment, // ++ on top of stack
    Token{representation:"assignTop"} : assignTop,

    // jumps
    Token{representation:"jmp"}: jmp,
    Token{representation:"jnz"}: jnz,


    // maths
    Token{representation:"add"}: add,

}

func exit(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    os.Exit(1)
}

func push(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    for i := 0; i < len(args); i++ {
        // try to resolve variable, otherwise push
        codeStack.Push(args[i])
    }
    (*ptr)++
}


func pop(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    codeStack.Pop()
    (*ptr)++
}

func decrement(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    topValue := codeStack.Pop()
    topValue = resolveToken(topValue)
    // just panic if someone tries to decrement a non-int
    topValue = topValue.(int) - 1
    codeStack.Push(topValue)
    (*ptr)++
}

func increment(codeStack *CodeStack, ptr *uint, args FunctionArgs){

}

func add(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    first := codeStack.Pop()
    second := codeStack.Pop()
    fValue := resolveToken(first)
    sValue := resolveToken(second)
    // these need to be ints, otherwise, you're screwed here. (Panics, but no auto-casting!)
    result := fValue.(int) + sValue.(int)
    codeStack.Push(result)
    (*ptr)++
}

/*
    Print the top of the stack
 */
func print(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    top := codeStack.Peek()
    top = resolveToken(top)
    t, instanceof := top.(int)
    var key string
    if instanceof {
        key = strconv.Itoa(t)
    } else {
        key = top.(string)
    }
    varValue, contains := variables[key]
    if contains{
        top = varValue.value
    }
    fmt.Print(top)
    (*ptr)++
}

/*
    Assign a value to a variable
 */
func assign(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    if len(args) < 2 {
        panic("Not enough arguments for variable assignment!")
    }
    varNameToken := args[0].(Token) // should be the variable name
    valueToken := resolveToken(args[1])
    variable := variables[varNameToken.representation]
    variable.assign(valueToken)
    variables[varNameToken.representation] = variable
    (*ptr)++
}

/*
    Assign the top of the stack to a variable
 */
func assignTop(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    if len(args) < 1{
        panic("not enough arguments!")
    }
    top := codeStack.Peek()
    tValue := resolveToken(top)
    arguments := make([]interface{},2)
    arguments[0] = args[0]
    arguments[1] = tValue
    assign(codeStack,ptr, arguments)
}

func jmp(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    if  len(args) != 1 {
        panic("Incorrect arguments passed to jmp statement!")
    }

    label := args[0].(Token).representation // label token starts with !, need to replace it with :
    label = ":" + label[1:]
    newptr := jumpTable[label]
    if newptr != 0 {
        *ptr = newptr
    } else {
        (*ptr)++
    }
}

func jnz(codeStack *CodeStack, ptr *uint, args FunctionArgs){
    top := codeStack.Peek()
    top = resolveToken(top)
    if top.(int) != 0 {
        label := args[0].(Token).representation // label token starts with !, need to replace it with :
        label = ":" + label[1:]
        newptr := jumpTable[label]
        *ptr = newptr
    }
    // else
    (*ptr)++
}


/**
Try to resolve the token as variable if possible, otherwise return the _token representation_ as value.
 */
func resolveToken(token interface{}) interface{}{

    actualToken, instanceof := token.(Token)
    if instanceof {
        variable, ok := variables[actualToken.representation]
        if ok {
            return variable.value
        }
        return actualToken.representation
    } else {
        return token // can't resolve it because it's not a token!
    }

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

func (t Token) String() string {
    return t.representation
}

type TokenLine struct{
    tokens []Token
    variableDeclaration bool               // Does this line contain variables, or code?
}


type Variable struct{
    datatype types.BasicKind
    value interface{}
}

func (v *Variable) assign(token interface{}){
    val := resolveToken(token)
    if v.datatype == types.String {
        v.value = val
    } else if v.datatype == types.Int {
        stringval, instanceof := val.(string)
        if instanceof{
            v.value,_ = strconv.Atoi(stringval)
        } else {
            v.value = val.(int)
        }
    }
}



/*
    Parse the incoming file, create an internal representations of the code.
    Since it is evaluated line by line, and we have goto statements, line numbers are important :-)
 */
func parse(source string){
    lines := strings.Split(source,"\n")
    lines = preParseFormat(lines)
    tokenLines := tokenize(lines)
    createJumpTable(tokenLines)

    stack := CodeStack{} // stack containing the values!


    instructionPointer := uint(0)
    for ;; {
        tokenLine := tokenLines[instructionPointer]
        if tokenLine.variableDeclaration {
            // create the variable map
            // each line should contain 2 tokens, the variable name + datatype
            name := tokenLine.tokens[0]
            dtype:= tokenLine.tokens[1]
            basicKind, defaultValue := typeFromString(dtype.representation)
            variable := Variable{datatype:basicKind, value:defaultValue}
            variables[name.representation] = variable
            instructionPointer++
            fmt.Println(variables)
        } else {
            // parse the code
            token := tokenLine.tokens[0]
            // skip labels
            if strings.HasPrefix(token.representation,":") {
                instructionPointer++
                continue
            }

            function := TokenFunctions[token]
            // maybe we have some args!
            args := make([]interface{},0)
            for i:= 1; i < len(tokenLine.tokens); i++{
                args = append(args, tokenLine.tokens[i])
            }
            function(&stack, &instructionPointer, args)
        }
    }
}

/*
    Create the indexes where we can jump through in the source code.
 */
func createJumpTable(tokenLines []TokenLine){
    // sweep the code, index the jumps!
    for i := 0; i < len(tokenLines); i++ {
        tokenLine := tokenLines[i]
        tokens := tokenLine.tokens
        for j := 0; j < len(tokens); j++ {
            token := tokens[j]
            if strings.HasPrefix(token.representation,":") {
                jumpTable[token.representation] = uint(i)
            }
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
        case "int":
            return types.Int, 0
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
            // parse instructions under .data
            // split the line based on ":"
            parts := strings.Split(line, ":")
            if len(parts) != 2{
                panic("Could not parse data section!")
            }
            variableName := Token{representation:strings.TrimSpace(parts[0])}
            variableType := Token{representation:strings.TrimSpace(parts[1])}
            tokeneLine := TokenLine{tokens:[]Token{variableName, variableType},variableDeclaration:true}
            tokenLines = append(tokenLines, tokenLine)
        } else {
            // parse instructions under .code
            parts := strings.Split(line," ")
            tokens := make([]Token,0)
            tokenizingString := false
            value := ""
            for i := 0; i < len(parts); i++ {
                part := parts[i]
                if strings.HasPrefix(part, "\"") {
                    // parsing a string
                    tokenizingString = true
                    value = part[1:]    // take after the "
                    continue // avoid adding this as a token, and just go to the next loop
                }
                if strings.HasSuffix(part, "\""){
                    value += part[:len(part)-1]
                    tokenizingString = false
                }
                if  tokenizingString{
                    value += part
                    continue
                }
                if value == "" {
                    value = part
                }
                token := Token{representation:strings.TrimSpace(value)}
                tokens = append(tokens, token)
                // reset value
                value = ""
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
    var quotationmark rune
    spacerune = 32
    quotationmark = 34
    inbetweenQuotes := false
    for i := 0; i < len(*source); i++ {
        line := (*source)[i]
        line = strings.Map(func(r rune) rune{
            if r == spacerune && !inbetweenQuotes{
                return -1
            }
            if r == quotationmark {
                inbetweenQuotes = !inbetweenQuotes
            }
            return r
        },line)
    }
}


/*
    Convenience function to print the source code (for testing)
 */
func printSource(source []string){
    for i := 0; i < len(source); i++ {
        fmt.Println(source[i])
    }
}