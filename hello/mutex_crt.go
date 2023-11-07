package main

import "time"
import "sync"

func main() {

    counter := 0
    var mu sync.Mutex
    for i:=0; i < 1000; i++{
        go func(){
            mu.Lock()
            defer mu.Unlock()
            counter += 1
        }()
    }
    time.Sleep(1* time.Second)
    mu.Lock()
    println("counter_value", counter)
    mu.Unlock()
}

