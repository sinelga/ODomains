package main

import (
	"domains"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"newdomainsV2/addnewdomain"
	"newdomainsV2/create_new_record"
	"os"
	"strings"
)

var token string

// The flag package provides a default help printer via -h switch

func main() {
	flag.Parse() // Scan the arguments list
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
		fmt.Println(token)


	}

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
			addnewdomain.Add(token, domaincsv)
			create_new_record.Create(token, domaincsv)
			

		}

	}

}
