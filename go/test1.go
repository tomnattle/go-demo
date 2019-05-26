package main
 
import (
    "fmt"
    "time"
    "os"
    "os/signal"
)
 
func main() {
    fmt.Println("定时器开始执行")
    go timer(timerFunc)
    c := make(chan os.Signal)
    signal.Notify(c)
    fmt.Println("启动")
    s := <-c
    fmt.Println("退出信号", s)
}

var count int
func timerFunc() {
    fmt.Printf("hello timer: %d\n", count)
    count++
}

func timer(timerfunc func()) {
    ticker := time.NewTicker(1*time.Second)
    for {
        select {
            case <-ticker.C:
                timerfunc()
        }
    }
}