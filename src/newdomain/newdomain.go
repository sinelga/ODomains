package newdomain

//https://api.digitalocean.com/domains/new?client_id=[your_client_id]&api_key=[your_api_key]&name=[domain]&ip_address=[ip_address]

//{"status":"OK","domain":{"id":117632,"name":"testsinelga.com"}}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"records"
	"domains"
	"extdomains"
)

type jsonresult struct {
	Status string `json:"status"`
	Domain domain `json:"domain"`
}

type domain struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

//func Add(your_client_id string, api_key string, domain string, ip_address string) {
  func Add(your_client_id string, api_key string, domain domains.Domaincsv) {		

	fmt.Println(domain.Name, domain.Ip)
	
	urlstr := "https://api.digitalocean.com/domains/new?client_id=" + your_client_id + "&api_key=" + api_key + "&name=" + domain.Name + "&ip_address=" + domain.Ip
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

		var data jsonresult
		err = json.Unmarshal(contents, &data)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Results: %v\n", data.Status)
		
		if data.Status=="OK" {
		
			records.AddCname(your_client_id,api_key,data.Domain.Id,"*","@")
			extdomains.AddToDB(domain)
		
		}		

	}
	

}
