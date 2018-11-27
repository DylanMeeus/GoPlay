// determine if there is a word pattern
package main

import (
    "fmt"
    "strings"
)

func main () {
    fmt.Printf("%v\n", wordPattern("abba","cat dog dog cat"))
}

func wordPattern(pattern string, str string) bool {
    patternMap := make(map[string]string, 0)
    ordered := make([]string,0)
    for _,part := range strings.Split(pattern, ""){
        ordered = append(ordered, part)
        patternMap[part] = ""
    }
    for _, word := range strings.Split(str, " ") {
        if _, ok := getKey(word, patternMap); !ok {
            // add to map at first free key.. but no iteration order.
            for _, o := range ordered {
                if patternMap[o] == "" {
                    patternMap[o] = word
                    break
                }
            }
        }
    }
    result := []string{}
    for _, o := range ordered {
        result = append(result, patternMap[o])
    }
    return strings.Join(result, " ") == str 
}

func getKey(value string, moap map[string]string) (key string, ok bool) {
    for k,v := range moap {
        if v == value {
            return k, true
        }
    }
    return "", false
}
