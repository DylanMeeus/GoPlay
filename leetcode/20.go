// valid parens
package main

import "fmt"


var charMapping = map[string]string {
    ")": "(",
    "}": "{",
    "]": "[",
}
func main() {
    fmt.Printf("%v\n",isValid("([])"))
}

func isValid(s string) bool {
    if s == "" {
        return true
    }
    parens := []string{}
    for i := 0; i < len(s); i++ {
        c := string(s[i])
        switch c {
        case "(":
            fallthrough
        case "[":
            fallthrough
        case "{":
            parens = append(parens,c)
            break
        default: 
            // it is not a starting token, so we delete
            if len(parens) == 0 {
                return false
            }
            if charMapping[c] == parens[len(parens)-1] {
                parens = parens[:len(parens)-1]
            } else {
                return false
            }
        }
    }
    fmt.Printf("%v\n", parens)
    return len(parens) == 0
}


