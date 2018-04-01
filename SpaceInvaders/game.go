package main

import(
    "runtime"
    "fmt"
    "github.com/go-gl/glfw/v3.2/glfw"
)


func init(){
    runtime.LockOSThread()
}

func main(){
    fmt.Println("Hello World")
    err := glfw.Init()
    if  err != nil {
        panic(err)
    }
    // end glfw at the end
    defer glfw.Terminate()

    window, err := glfw.CreateWindow(640,480, "Testing", nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()

    for !window.ShouldClose(){
        window.SwapBuffers()
        glfw.PollEvents()
    }
}
