package main

import (
	"fmt"
	"time"
)

func executeA(c chan string) {
    for {
        time.Sleep(time.Millisecond * 500)
        c <- "Canal 1"
    }
}

func executeB(c chan string) {
    for {
        time.Sleep(time.Second * 2)
        c <- "Canal 2"
    }
}

func main() {
    channelA, channelB := make(chan string), make(chan string)

    go executeA(channelA)
    go executeB(channelB)

    // go func() {
    //     for {
    //         time.Sleep(time.Second * 2)
    //         channelB <- "Canal 2"
    //     }
    // }()

    for {
        // msgA := <- channelA
        // fmt.Println(msgA)

        // msgB := <- channelB
        // fmt.Println(msgB)

        select {
        case msgA := <- channelA:
            fmt.Println(msgA)
        case msgB := <- channelB:
            fmt.Println(msgB)
        }
    }
}