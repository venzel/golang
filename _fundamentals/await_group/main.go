package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(3)

    go callA(&wg)
    go callB(&wg)
    go callC(&wg)

    wg.Wait()
}

func callA(wg *sync.WaitGroup) {
    time.Sleep(time.Second)
    fmt.Println("Finalizou A")
    wg.Done()
}

func callB(wg *sync.WaitGroup) {
    time.Sleep(time.Second)
    fmt.Println("Finalizou B")
    wg.Done()
}

func callC(wg *sync.WaitGroup) {
    time.Sleep(time.Second)
    fmt.Println("Finalizou C")
    wg.Done()
}