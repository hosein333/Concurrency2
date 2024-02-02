package main

import (
	"fmt"
	"time"
)

func printLog(pChannel chan string) {
    for {
        v, ok := <-pChannel

        if !ok {
            fmt.Println("pChannel Closed")
            return
        }

        fmt.Println(v)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    LogChannel := make(chan string, 10)

    go printLog(LogChannel)

    for i := 0; i < 10; i++ {
        LogChannel <- fmt.Sprintf("Coount is %d",i)
    }
    close(LogChannel)

    time.Sleep(10 * time.Second)

}
