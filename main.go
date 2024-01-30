package main

import (
	"fmt"
	"time"
)

func printLog(pChannel chan string) {
    for i := 0; i < 10; i++ {
        v := <-pChannel
        fmt.Println(v)
    }
    fmt.Println("exit printLog Function")
}

func main() {
    LogChannel := make(chan string, 10)

    go printLog(LogChannel)

    for i := 0; i < 30; i++ {
        LogChannel <- fmt.Sprintf("Coount is %d",i)

    }

    fmt.Println("exit fot loop")
    
    time.Sleep(10 * time.Second)

}
