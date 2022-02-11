package main

import "bhg-scanner/scanner"

func main(){
	scanner.PortScanner(0,1024,true)//checks all ports, displays results
}