package main

import "fmt"
import "strings"

func main() {
    fmt.Printf("%v\n", reverseVowels("leetcode"))
}

func isVowel(s *string) bool {
    vowels := []string{"a","e","i","o","u", "A", "E", "I", "O", "U"}
    for _, v := range(vowels) {
        if *s == v {
            return true
        }
    }
    return false
}

// we create an array with gaps in it for the vowels
// then we fill in the vowels in reverse order of collections
func reverseVowels(s string) string {
    parts := strings.Split(s, "")
    vowels := make([]string, len(parts)) 
    vi := 0
    buff := make([]string, len(parts))
    for index, letter := range(parts){
        if isVowel(&letter) {
            vowels[vi] = letter
            vi++
        } else {
            buff[index] = letter
        }
    }
    fmt.Println(vowels)
    fmt.Println(buff)
    for i, l := range(buff) {
        if l == "" {
            buff[i] = vowels[vi-1]
            // empty vowel buffer
            vi--
        }
    }
    return strings.Join(buff, "") 
}
