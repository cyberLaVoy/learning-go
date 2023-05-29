package main

import (
    "fmt"
    "math/rand"
    "time"
)

func fanInOrd(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            c <- <-input1
        }
    }()
    go func() {
        for {
            c <- <-input2
        }
    }()
    return c
}

func fanInSel(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case s := <-input1:
                c <- s
            case s := <-input2:
                c <- s
            }
        }
    }()
    return c
}

func boring(message string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            c <- message
            time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
        }
    }()
    return c
}

func boringWithQuit(message string, quit chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case c <- message:
                time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
            case s := <-quit:
                //clean up
                fmt.Println(s)
                quit <- "See yah!"
                return
            }
        }
    }()
    return c
}

func main() {
    quit := make(chan string)
    c := boringWithQuit("Give me the order to stop...", quit)
    for i := 0; i < 10; i++ {
        fmt.Println(<-c)
    }
    quit <- "Okay, stop."
    fmt.Println(<-quit)

    //c := fanInOrd(boring("Boring 1..."), boring("Boring 2..."))
    ca := fanInSel(boring("Boring 1..."), boring("Boring 2..."))
    for i := 0; i < 10; i++ {
        fmt.Println(<-ca)
    }
}
