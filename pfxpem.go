package main

import (
	"encoding/pem"
	"flag"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"os"
)

var source *string = flag.String("source", "", "Source PFX file")
var password *string = flag.String("password", "", "Password for PFX file")

func main() {
	flag.Parse()

	srcFilePath := *source

	if srcFilePath == "" {
		fmt.Println("Argument 'source' required. Please see information with -h.")
		os.Exit(2)
		return
	}

	// Buffer file

	data, err := ioutil.ReadFile(srcFilePath)

	if err != nil {
		panic(err)
	}

	// Convert PFX to multiple PEM.Block

	pemBlocks, err := pkcs12.ToPEM(data, *password)

	if err != nil {
		panic(err)
	}

	// Display PEM

	for i, block := range pemBlocks {
		fmt.Println("PEM Block", i+1, ":")
		fmt.Println(string(pem.EncodeToMemory(block)))
	}
}
