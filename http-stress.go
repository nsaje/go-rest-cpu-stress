package main

import (
    "fmt"
    "net/http"
    "time"
)

var burnCpu = make(chan bool)

func cpuBurn() {
    var burn bool = false
    for {
        select {
        case burn = <-burnCpu:
        default:
            if burn {
                _ = 5*5
            } else {
                time.Sleep(100 * time.Millisecond)
            }
        }
    }
}

func startCpu(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Starting CPU burn\n")
    burnCpu <- true
}

func stopCpu(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Stopping CPU burn\n")
    burnCpu <- false
}

func main() {
    go cpuBurn()
    http.HandleFunc("/start_cpu", startCpu)
    http.HandleFunc("/stop_cpu", stopCpu)
    http.ListenAndServe(":8080", nil)
}
