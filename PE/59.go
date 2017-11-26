package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "os"
)

func main(){
    characters := getInputChars()
    tryDecrypt(characters)
}

/*
    Get possible decryptions
 */
func tryDecrypt(characters []string){
    fmt.Println("Start decryption")
    // combine characters from 'a' to 'z' in ASCII
    start := 97
    limit := 122
    for a := start; a <= limit; a++  {
        for b := start; b <= limit; b++{
            for c := start; c <= limit; c++{
                keyArr := []int{a,b,c}
                decryptedAscii := decrypt(characters, keyArr)
                decryptedText := convertAsciiString(decryptedAscii)
                if isEnglish(decryptedText) {
                    // sum the ascii values
                    fmt.Println(decryptedText)
                    fmt.Println(sumAscii(decryptedAscii))
                    os.Exit(1)
                }
            }
        }
    }
}

func sumAscii(characters []int) uint64{
    var sum uint64
    sum = 0
    for i := 0; i < len(characters); i++ {
        fmt.Println(characters[i])
        asciivalue := uint64(characters[i])
        sum += asciivalue
    }
    return sum
}

/*
    Is this text english?.
 */
func isEnglish(text string) bool{
    // english usually contains "I" as a common word
    words := strings.Split(text," ")
    if len(words) == 1 { //
        return false
    }

    that := false
    for i := 0; i < len(words); i++ {
        if words[i] == "that" {
            that = true
        }

        if that {
            return true
        }
    }
    return false
}

func decrypt(characters []string, asciiKeyCodes []int) []int{
    decryptedAscii := make([]int, len(characters))
    for i := 0; i < len(characters); i++ {
        // string to int
        character := characters[i]
        if len(character) > 2 && strings.HasSuffix(character, "\n") {
            character = character[0:len(character)-2]
        }
        value, err := strconv.Atoi(character)
        if err != nil {
            panic(err)
        }
        decrypted := value ^ (asciiKeyCodes[i % 3])
        decryptedAscii = append(decryptedAscii, decrypted)
    }
    return decryptedAscii
}

func convertAsciiString(ascii []int) string{
    chars := make([]string,len(ascii))
    for i := 0; i < len(ascii); i++ {
        r := rune(ascii[i])
        chars = append(chars, string(r))
    }
    return strings.Join(chars, "")
}

func getInputChars() []string{
    input, err := ioutil.ReadFile("resources/p059_cipher.txt")

    if  err != nil {
        panic(err)
    }
    chars := strings.Split(string(input), ",")
    return chars
}