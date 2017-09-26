package main


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