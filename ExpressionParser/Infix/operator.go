package InfixParser


var operatormap = map[string] Operator{
    "+": Operator{symbol:"+", precedence:2, leftAssociative:true},
    "-": Operator{symbol:"-", precedence:2, leftAssociative:true},
    "/": Operator{symbol:"/", precedence:3, leftAssociative:true},
    "*": Operator{symbol:"*", precedence:3, leftAssociative:true},
}

type Operator struct{
    symbol string
    precedence int
    leftAssociative bool // it's either left or right.. only need a bit now
}


func OperatorFromString(s string) Operator{
    return operatormap[s]
}

func (o Operator) String() string{
    return o.symbol
}

func (o Operator) HasPrecedence(other Operator) bool{
    return o.precedence > other.precedence
}

func (o Operator) calculate(a float64, b float64) float64 {
    switch o{
        case operatormap["+"]:
            return add(a,b)
        case operatormap["-"]:
            return subtract(a,b)
        case operatormap["*"]:
            return multi(a,b)
        case operatormap["/"]:
            return div(a,b)
    }
    return 0 // this should probably not be zero
}

func add(a float64, b float64) float64 {
    return a + b
}

func subtract(a float64, b float64) float64 {
    return a - b
}

func multi(a float64, b float64) float64{
    return a * b
}

func div(a float64, b float64) float64 {
    return a / b
}