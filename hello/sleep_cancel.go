package main

import "time"
import "sync"

var done bool
var mu sync.Mutex

func main() {

    println("let us start this one")
    go Periodic()
    time.Sleep(5 * time.Second)
    mu.Lock()
    done = true
    mu.Unlock()
    println("cancelled")
    time.Sleep(2 * time.Second)


}

func Periodic(){
    for {
        mu.Lock()
        if done{
            mu.Unlock()
            return

        }
        mu.Unlock()
        println("tick")
        time.Sleep(1 * time.Second)
    }

}
