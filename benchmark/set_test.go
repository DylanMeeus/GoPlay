package main

import (
    "testing"
)

var (
    empty = struct{}{}
)

func BenchmarkBoolSet(b *testing.B) {
    m := make(map[int]bool,0)
    for i := 0; i < b.N; i++ {
        m[i] = true
    }
}

func BenchmarkStructSet(b *testing.B) {
    m := make(map[int]struct{},0)
    for i := 0; i < b.N; i++ {
        m[i] = struct{}{}
    }
}

func BenchmarkStructOutside(b *testing.B) {
    m := make(map[int]struct{},0)
    for i := 0; i < b.N; i++ {
        m[i] = empty 
    }
}

func BenchmarkInterfaceSet(b *testing.B) {
    m := make(map[int]interface{})
    for i := 0; i < b.N; i++ {
        m[i] = struct{}{}
    }
}

func BenchmarkPointerSet(b *testing.B) {
    m := make(map[int]*struct{},0)
    for i := 0; i < b.N; i++ {
        m[i] = &empty 
    }
}
