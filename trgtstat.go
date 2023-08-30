package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var pingTargets = []string{
    "IP ADDRESS",
    "DOMAIN NAME",
} 

func main() { 
    for _, ip := range pingTargets {
        if ipPing(ip) {
            fmt.Printf("\u2713 %s is reachable\n", ip)
        } else {
            fmt.Printf("\u2717 %s is unreachable\n", ip)   
        }
        traceRoute(ip)
    }
}

func ipPing(ip string) bool {
    fmt.Println(strings.Repeat("_", 80))
    cmd := exec.Command("ping", "-c", "1", ip)
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Failed to execute ping: ", err)
        return false
    }
    
    fmt.Println(string(output))
    fmt.Println()
    return true
}

func traceRoute(ip string) bool {
    fmt.Println(strings.Repeat("_", 80))
    fmt.Printf("Executing traceroute on target...\n")
    cmd := exec.Command("traceroute", ip)

    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Failed to execute traceroute: ", err)
        return false
    }
    
    fmt.Println(string(output))
    fmt.Println()
    return true
}
