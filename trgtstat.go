package main

import (
	"fmt"
	"os/exec"
        "strings"
)

var pingTargets = []string{  //Input targets in this array
    "IP ADDRESS",
    "DOMAIN NAME",
} 

func main() { 
    var choice int
    for {
        fmt.Println("--- Target Status ---")
        fmt.Println("1. Ping targets")
        fmt.Println("2. Traceroute")
        fmt.Println("3. Nmap scan")
        fmt.Println("4. All")
        fmt.Println("5. Exit")
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
                if traceRoute(ip) {
                    fmt.Printf("\u2713 %s is reachable\n", ip)
                } else {
                    fmt.Printf("\u2717 %s is unreachable\n", ip)   
                }
            }
        case 3:
            for _, ip := range pingTargets {
                nmapScan(ip)
            }
        case 4:
            for _, ip := range pingTargets {
                if ipPing(ip) {
                    fmt.Printf("\u2713 %s is reachable\n", ip)
                } else {
                    fmt.Printf("\u2717 %s is unreachable\n", ip)   
                }
            traceRoute(ip)
            nmapScan(ip)
            }
        case 5:
            fmt.Println("Exiting...")
            return
        }
        break
    }
}

func ipPing(ip string) bool {
    fmt.Println(strings.Repeat("_", 80))
    fmt.Println("Pinging target...")
    cmd := exec.Command("ping", "-c", "4", ip)
    
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


func nmapScan(ip string) bool {
    fmt.Println(strings.Repeat("_", 80))
    fmt.Printf("Running nmap on target...\n")
    cmd := exec.Command("nmap", "-p 0-1000", "-Pn", ip)

    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Failed to execute nmap scan: ", err)
        return false
    }
    
    fmt.Println(string(output))
    fmt.Println()
    return true
}
