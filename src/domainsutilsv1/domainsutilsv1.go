package main 

import (
    "flag"
    "fmt"
    "encoding/csv"
    "os"
    "io"
    "net/http"
    "io/ioutil"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")


var your_client_id string
var api_key string
var domain string
var ip_address string


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
	fmt.Println("id:",your_client_id, api_key)

	csvFiled, err := os.Open("domains.csv")
	defer csvFiled.Close()
	if err != nil {
		panic(err)
	}

   	urlstr := "https://api.digitalocean.com/v1/domains/k18.me/records?client_id=" + your_client_id + "&api_key=" + api_key
	fmt.Println(urlstr)
	
		response, err := http.Get(urlstr)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))

    }
        
    
}

