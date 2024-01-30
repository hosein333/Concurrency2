package main

import (
	"fmt"
	"time"
)

func printLog(pChannel chan string) {
    for {
        v := <-pChannel
        fmt.Println()
        fmt.Println(v)
        time.Sleep(1 * time.Second)
        
    }
}

func main() {
    LogChannel := make(chan string, 10)

    go printLog(LogChannel)

    for i := 0; i < 20; i++ {
        fmt.Printf("\nWriting Count %d", i)
        LogChannel <- fmt.Sprintf("Coount is %d",i)
        fmt.Printf("\nWited Count %d", i)
    }

    fmt.Println("exit fot loop")
    
    time.Sleep(1 * time.Second)

}
