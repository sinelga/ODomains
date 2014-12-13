package main 

import (
    "flag"
    "fmt"
   "encoding/csv"
   "os"
    "io"
    "getalldomains"
)

const APP_VERSION = "0.1"

var your_client_id string
var api_key string

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
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
	fmt.Println(your_client_id, api_key)
	
	getalldomains.GetAllD(your_client_id,api_key)
    
    
}

