package main

import (
	"fmt"
	"time"
)

const (
	DOUBLE_CLICK_LIMIT_MS = 100
)

var (
	lastClicked = map[interface{}]int64{}
)

func main() {

	go func() {
		cleanupTimer := time.After(10 * time.Millisecond)
		for {
			select {
			case <-cleanupTimer:
				cleanup()
			}
		}
	}()

	obj := struct{}{} // some reference object
	fmt.Printf("%v\n", IsDoubleClick(obj))
	fmt.Printf("%v\n", IsDoubleClick(obj))
	fmt.Printf("%v\n", IsDoubleClick(obj))
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%v\n", IsDoubleClick(obj))
	fmt.Printf("%v\n", lastClicked)
	cleanup()
	time.Sleep(4000 * time.Millisecond)
	cleanup()
	fmt.Printf("%v\n", lastClicked)
}

// IsDoubleClick checks if an object is clicked on twice
func IsDoubleClick(i interface{}) bool {
	if i == nil {
		return false
	}

	if lastTime, ok := lastClicked[i]; ok && time.Now().UnixMilli()-lastTime < DOUBLE_CLICK_LIMIT_MS {
		fmt.Printf("diff: %v and limit :%v\n", time.Now().UnixMilli()-lastTime, DOUBLE_CLICK_LIMIT_MS)
		delete(lastClicked, i)
		return true
	}

	lastClicked[i] = time.Now().UnixMilli()
	//cleanup()
	return false
}

func cleanup() {
	fmt.Println("cleaning up")
	for k, lastTime := range lastClicked {
		if lastTime+DOUBLE_CLICK_LIMIT_MS < time.Now().UnixMilli() {
			delete(lastClicked, k)
		}
	}
}
