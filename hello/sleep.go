package main

import "time"

func main() {
    println("let's start this ")
    time.Sleep(1 * time.Second)
    go Periodic()
    time.Sleep(5 * time.Second)
}


func Periodic(){
    for {
        println("tick")
        time.Sleep(1 * time.Second)
    }

}
