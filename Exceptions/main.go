package main

import (
    "fmt"
    "errors"
)

type errorfunc func() error

func main() {
    var err error
    if  ((err = w() ||
        err = x())); err != nil {
        panic(err)
    }
}

func errorable(fs ...errorfunc) error {
    for i,f := range fs { 
        e := f()
        if e != nil {
            return errors.New(fmt.Sprintf("Failed at function: %v, iteration: %v\n", f, i))
        }
    }
    return nil
}

func w() error {
    fmt.Println("w")
    return nil
}
func x() error {
    fmt.Println("x")
    return nil
}
func y() error {
    fmt.Println("y")
    return errors.New("Failed to print y") 
}
func z() error {
    fmt.Println("z")
    return nil
}
