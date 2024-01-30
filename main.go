package main

import (
	"fmt"
	"time"
)

func printLog(pChannel chan string) {
    for {
        v := <-pChannel
        fmt.Println(v)
    }
}

func main() {
    LogChannel := make(chan string, 1000)

    go printLog(LogChannel)

    for i := 0; i < 1000; i++ {
        LogChannel <- fmt.Sprintf("Coount is %d",i)

    }
    time.Sleep(10 * time.Second)

}
