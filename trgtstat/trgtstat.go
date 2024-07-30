package main

import (
	"fmt"
	"os/exec"
    	"os"
	"strings"
	"encoding/json"
	)

// Input all targets (IPs & domains) into the targets.json file.
type Config struct {
    PingTargets []string `json:"pingTargets"`
}

func main() { 
    var choice int

    file, err := os.Open("targets.json")
    if err != nil {
        fmt.Println("Error opening file: ", err)
        return
    }
    defer file.Close()

    config := Config{}
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
            fmt.Println("Error decoding JSON: ", err)
            return
    }

    pingTargets := config.PingTargets

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
            allReachable := true
            var unreachableTargets []string

            for _, ip := range pingTargets {
                if ipPing(ip) {
                    fmt.Printf("\u2713 %s is reachable\n", ip)
                } else {
                    fmt.Printf("\u2717 %s is unreachable\n", ip)   
                    allReachable = false
                    unreachableTargets = append(unreachableTargets, ip)
                }
            }

                if allReachable {
                    fmt.Println("\n\u2713 All targets are reachable.")
                } else {
                    fmt.Println("\n\u2717 Not all targets are reachable. Unreachable targets:", strings.Join(unreachableTargets, ", "))
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
    cmd := exec.Command("ping", "-c", "2", ip)
    
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
