package main

import (
	"fmt"
	"github.com/DylanMeeus/GoPlay/ExpressionParser/Infix"

	"io/ioutil"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	out := 0
	for _, line := range strings.Split(string(in), "\n") {
		if line == "" {
			continue
		}
		out += int(InfixParser.Eval(line))
	}

	fmt.Println(out)
}
