// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// From course-materials/materials/lab/2/bhg-scanner/main
// go build; ./main


package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)



func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.DialTimeout("tcp", address,750*time.Millisecond)
		if err != nil { 
			results <- -1*p
			continue
		}
		conn.Close()
		results <- p
	}
}


func PortScanner(lo int, hi int, verbose bool) (int, int, map[int]bool){  
	var openports []int
	var closedports []int
	portlist := make(map[int]bool)
	ports := make(chan int, 100)
	results := make(chan int)

	for i := lo; i <= hi; i++ {
		go worker(ports, results)
	}

	go func() {
		for i := lo+1; i <= hi; i++ {
			ports <- i
		}
	}()
	for i := lo; i < hi; i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
			portlist[port]=true
		}else if port < 0{
			closedports = append(closedports, -1*port)
			portlist[port]=false
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)
	if verbose{
		for _, port := range openports {
			fmt.Printf("%d open\n", port)
		}
	
		for _, port := range closedports {
			fmt.Printf("%d closed\n", port)
		}
	}

	return len(openports), len(closedports), portlist
}
