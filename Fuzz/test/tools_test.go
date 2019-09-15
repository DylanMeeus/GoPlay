package tools

import "testing"
import "../tools"

func TestHello(t *testing.T) {
    res := tools.Hello("Dylan")
    if res != "Hello: Dylan" {
        t.Errorf("Incorrect output, got %v but expected %v", res, "Hello: Dylan")
    }
}

func Fuzz(data []byte) int {
    tools.Hello(string(data))
    return 0
}
