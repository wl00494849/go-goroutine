package controller

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var counter = 0

func Goroutine1(ctx *gin.Context) {
	var msg = "hellow"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)

	msg = "goodbyg"
}

func Gorutine2(ctx *gin.Context) {
	thread := runtime.GOMAXPROCS(-1)
	// -1 = max = 8
	fmt.Printf("Threads: %v\n", thread)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sey()
		m.Lock()
		go increase()
	}
	wg.Wait()
}

func sey() {
	fmt.Printf("Test %v\n", counter)
	m.RUnlock()
	time.Sleep(time.Second)
	wg.Done()
}

func increase() {
	counter++
	m.Unlock()
	wg.Done()
}
