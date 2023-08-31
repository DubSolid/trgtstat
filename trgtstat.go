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
    var choice int
    for {
        fmt.Println("--- Target Status ---")
        fmt.Println("1. Ping targets")
        fmt.Println("2. Traceroute on targets")
        fmt.Println("3. Both")
        fmt.Println("4. Exit")
	fmt.Print("Choose and option: ")

        fmt.Scanln(&choice)
        switch choice {
        case 1:      
            for _, ip := range pingTargets {
                if ipPing(ip) {
                    fmt.Printf("\u2713 %s is reachable\n", ip)
                } else {
                    fmt.Printf("\u2717 %s is unreachable\n", ip)   
                }
            }
        case 2:
            for _, ip := range pingTargets {
                traceRouteOnly(ip)
            }
        case 3:
            for _, ip := range pingTargets {
                if ipPing(ip) {
                    fmt.Printf("\u2713 %s is reachable\n", ip)
                } else {
                    fmt.Printf("\u2717 %s is unreachable\n", ip)   
                }
            traceRoute(ip)
            }
        case 4:
            fmt.Println("Exitingg...")
            return

        }
        break
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

func traceRouteOnly(ip string) bool {
    fmt.Println(strings.Repeat("_", 80))
    fmt.Printf("Executing traceroute on selected targets...\n")
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
