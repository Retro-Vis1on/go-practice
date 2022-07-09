package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX,hasSPF,SPFRecords,hasDMARC,DMARCRecords")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFRecords, DMARCRecords string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	} else {
		hasMX = false
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, rec := range txtRecords {
		if strings.HasPrefix(rec, "v=spf1") {
			hasSPF = true
			SPFRecords = rec
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, rec := range dmarcRecords {
		if strings.HasPrefix(rec, "v=DMARC1") {
			hasDMARC = true
			DMARCRecords = rec
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, SPFRecords, hasDMARC, DMARCRecords)
}
