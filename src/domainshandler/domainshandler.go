package main

import (
	"domains"
	"encoding/csv"
	"fmt"
	"io"
	"newdomain"
	"os"
	"strings"
)

var your_client_id string
var api_key string
var domain string
var ip_address string


func main() {

	csvFile, err := os.Open("credapi.csv")
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

		your_client_id = fields[0]
		api_key = fields[1]

	}
	fmt.Println("id:",your_client_id, api_key)

	csvFiled, err := os.Open("domains.csv")
	defer csvFiled.Close()
	if err != nil {
		panic(err)
	}



	csvReaderd := csv.NewReader(csvFiled)
	csvReaderd.TrailingComma = true

	var domaincsv domains.Domaincsv

	for {
		fields, err := csvReaderd.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if !strings.HasPrefix(fields[0], "#") {

			domaincsv.Locale = fields[0]
			domaincsv.Themes = fields[1]
			domaincsv.Name = fields[2]
			domaincsv.Ip = fields[3]

			newdomain.Add(your_client_id, api_key, domaincsv)

		}

	}

}
