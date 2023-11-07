package main

import "time"

func main() {

    counter := 1
    for i:=0; i < 1000; i++{
        go func(){
            counter += 1
        }()
    }
    time.Sleep(1* time.Second)
    println("counter_value", counter)
}

