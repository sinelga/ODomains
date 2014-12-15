package main

import (
	"changeip"
	"dumpintocsv"
	"encoding/csv"
	"flag"
	"fmt"
	"getalldomainsv2"
	"io"
	"os"
)

const APP_VERSION = "0.1"

var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var fromIpFlag *string = flag.String("fromIp", "", "must be valid ip address")

var token string

func main() {
	flag.Parse() // Scan the arguments list

	var fromIp = *fromIpFlag

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	csvFile, err := os.Open("token.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(csvFile)

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		token = fields[0]

	}

	var recordstodump [][]string

	alldomainrecords := getalldomainsv2.GetAllD(token, "https://api.digitalocean.com/v2/domains")

	for _, domainrecords := range alldomainrecords {

		if domainrecords[1] == fromIp {

			recordtodump := changeip.ChangeFromTo(token, domainrecords[0], domainrecords[1], "104.131.95.7")
			//			recordtodump :=[]string{domainrecords[0],domainrecords[1],"104.131.95.7",}
			recordstodump = append(recordstodump, recordtodump)

		}

	}
	dumpintocsv.Dump(recordstodump)
}
