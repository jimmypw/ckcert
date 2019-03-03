package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	exitOK      = 0
	exitRENEW   = 1
	exitEXPIRED = 2
	exitERROR   = 3
)

func main() {
	currenttime := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("No input file specified")
		os.Exit(exitERROR)
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: Unable to read certificate.")
		os.Exit(exitERROR)
	}
	decoded, rest := pem.Decode(file)
	_ = rest
	parsedcertificate, err := x509.ParseCertificate(decoded.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(exitERROR)
	}
	validity := parsedcertificate.NotAfter.Sub(parsedcertificate.NotBefore)
	remaining := parsedcertificate.NotAfter.Sub(currenttime)

	if remaining.Hours() < 0 {
		fmt.Println("The certificate has expired.")
		os.Exit(exitEXPIRED)
	}

	if remaining.Hours() < (validity.Hours() / 2) {
		fmt.Println("The certificate is approaching expiry.")
		os.Exit(exitRENEW)
	}

	fmt.Println("The certificate is OK!")
	os.Exit(exitOK)
}
