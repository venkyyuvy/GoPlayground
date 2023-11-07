package main

import "sync"

func main() {
    var wait_var sync.WaitGroup;
    for i := 0; i < 5; i++ {
        wait_var.Add(1)
        go func(){
            defer wait_var.Done()
            sendRPC(i)
        }()
    }
    wait_var.Wait()
}

func sendRPC(i int){
    println(i)
}

