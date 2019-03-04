package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	exitOK      = 0
	exitRENEW   = 1
	exitERROR   = 2
	modeDAYS    = 3
	modePERCENT = 4
)

// Config struct will hold command line switch configuration.
type Config struct {
	mode int
	days *int
	prct *int
	file *string
}

func isFlagSet(name string) bool {
	isset := false
	flag.Visit(func(thisflag *flag.Flag) {
		if thisflag.Name == name {
			isset = true
		}
	})

	return isset
}

func processFlags() *Config {
	conf := new(Config)
	conf.file = flag.String("f", "", "The certificate file")
	conf.days = flag.Int("d", 0, "Number of valid days to check for on certificate")
	conf.prct = flag.Int("p", 0, "Percent of relative validity remaining on the certificate")
	flag.Parse()

	if !isFlagSet("f") {
		fmt.Println("No input file (-f) specified")
		os.Exit(exitERROR)
	}

	if isFlagSet("d") && isFlagSet("p") {
		fmt.Println("Switches -p and -d are mutrally exclusive")
		os.Exit(exitERROR)
	}

	if (!isFlagSet("d")) && (!isFlagSet("p")) {
		fmt.Println("Either -p or -d must be specified")
		os.Exit(exitERROR)
	}

	if isFlagSet("d") {
		if *conf.days < 1 {
			fmt.Println("-d must be a positive integer")
			os.Exit(exitERROR)
		}
		conf.mode = modeDAYS
	} else {
		if (*conf.prct < 1) || (*conf.prct > 100) {
			fmt.Println("-p must be a positive integer between 1 and 100")
			os.Exit(exitERROR)
		}
		conf.mode = modePERCENT
	}

	return conf
}

func readAndParseCertificate(certpath string) *x509.Certificate {
	file, err := ioutil.ReadFile(certpath)
	if err != nil {
		fmt.Println("ERROR: Unable to read certificate.")
		os.Exit(exitERROR)
	}
	decoded, _ := pem.Decode(file)
	parsedcertificate, err := x509.ParseCertificate(decoded.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(exitERROR)
	}
	return parsedcertificate
}

func hoursToDays(hours float64) float64 {
	return hours / 24.0
}

func certExpiring() {
	fmt.Println("The certificate is approaching expiry.")
	os.Exit(exitRENEW)
}
func certOk() {
	fmt.Println("The certificate is OK!")
	os.Exit(exitOK)
}

func daysRemaining(begin, end time.Time, minimumdays int) bool {
	returndata := true
	delta := end.Sub(begin)
	if hoursToDays(delta.Hours()) < float64(minimumdays) {
		// Certificate is exiring
		returndata = true
	} else {
		// Certificate is not expiring
		returndata = false
	}

	return returndata
}

func percentRemaining(thetime time.Time, cert *x509.Certificate, percent int) bool {
	returndata := true
	validity := cert.NotAfter.Sub(cert.NotBefore)
	minimumvalidity := (validity.Hours() / 100.0) * float64(percent)
	deltaremaining := cert.NotAfter.Sub(thetime)
	if deltaremaining.Hours() < minimumvalidity {
		// Certificate is exiring
		returndata = true
	} else {
		// Certificate is not expiring
		returndata = false
	}
	return returndata
}

func main() {
	config := processFlags()
	cert := readAndParseCertificate(*config.file)

	currenttime := time.Now()

	switch config.mode {
	case modeDAYS:
		if daysRemaining(currenttime, cert.NotAfter, *config.days) {
			certExpiring()
		} else {
			certOk()
		}
	case modePERCENT:
		if percentRemaining(currenttime, cert, *config.prct) {
			certExpiring()
		} else {
			certOk()
		}
	}
}
