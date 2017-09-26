// Infix parser using the Shunting-yard algorithm

package main

import (
    "fmt"
    "strings"
    "../Collections"
    "strconv"
)

func main(){
    eval("13+250 - 2")
}


func eval(equation string){
    // I) Apply Shunting-Yard (returns Reverse Polish Notation of equation)
    // II) Solve Reverse Polish Notation
    fmt.Println("Parsing: " + equation)
    eqparts := ShuntingYard(equation)
    solveRPN(eqparts)
}

// remove some of the spaces - that is the only cleanup at the moment
func formatInputString(input string) string{
    var spacerune rune
    spacerune = 32
    return strings.Map(func(r rune) rune{
        if r == spacerune{
            return -1
        }
        return r
    },input)
}

func tokenize(input string) []string{
    tokens := []string{}
    index := 0
    for index < len(input) {
        token := ""
        charvalue := input[index]
        for charIsDigit(charvalue) && index+1 < len(input){ // also make sure it is in the bounds
            token += string(charvalue)
            index++
            charvalue = input[index]
        }

        if token == ""{
            token = string(charvalue)
            index++
        }
        tokens = append(tokens, token)
    }
    return tokens
}

// evaluate a string, returns the Reverse Polish Notation..
func ShuntingYard(input string) []interface{}{
    input = formatInputString(input)
    tokens := tokenize(input)

    output := make([]interface{},0)
    operators := collections.Stack{}

    for i := 0; i < len(tokens); i++{
        token := tokens[i]
        if isNumber(token){
            output = append(output,token)
        } else { // it is an operator

            operator := OperatorFromString(token)
            operators = append(operators,operator)
        }
    }

    // now pop the operators onto output
    for operators.Empty() {
        top := operators.Pop()
        output = append(output,top)
    }

    return output
}

func isNumber(token string) bool {
    // if the first digit is a number, it is a number
    return charIsDigit(token[0])
}

func charIsDigit(char byte) bool{
    zero := "0"[0]
    nine := "9"[0]
    return char >= zero && char <= nine
}


func solveRPN(eq []interface{}) int{
    fmt.Println(eq)

    // we keep stacking it.. and applying an operator to the stack when we encounter it, and shrink, etc..

    execstack := collections.Stack{}
    for i := 0; i < len(eq); i++{
        token := eq[i]
        operator, isOperator := token.(Operator)
        number, isNumber := token.(string)
        if isOperator{
            as := execstack.Pop().(string)
            bs := execstack.Pop().(string)
            a, erra := strconv.ParseFloat(as,64)
            b, errb := strconv.ParseFloat(bs,64)
            if erra != nil {
                panic(erra)
            }
            if errb != nil {
                panic(errb)
            }
            calcresult := operator.calculate(a,b)
            calcstring := strconv.FormatFloat(calcresult, 'f', -1, 64)
            execstack.Push(calcstring)
        } else if isNumber{
            execstack.Push(number)
        }
    }
    fmt.Println(execstack)
    return 0
}