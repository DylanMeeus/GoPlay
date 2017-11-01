package main


// main package for alewig


import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

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

/*
    Parse the incoming file, create an internal representations of the code.
    Since it is evaluated line by line, and we have goto statements, line numbers are important :-)
 */
func parse(source string){
    lines := strings.Split(source,"\n")
    lines = preParseFormat(lines)
}

/*
    Format the code before parsing to remove some unneccessary elements
    e.g: No empty lines, comments, ..
 */
func preParseFormat(source []string) []string{ // keep this side-effect free, take the perf hit.

    //formatted := make([]string,len(source))
    deleteComments(&source)
    deleteEmptyLines(&source)
    printSource(source)

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

func printSource(source []string){
    for i := 0; i < len(source); i++ {
        fmt.Println(source[i])
    }
}