package httpgetter

import (
	"bytes"
	"domains"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAll(token string, urlstr string) []byte {

	var retcontents []byte

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlstr, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	//	req.Header.Add("Content-Type","application/json")

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		retcontents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

	}

	return retcontents

}

func PostNewDomain(token string, newdomain domains.NewDomain) []byte {

	var retcontents []byte

	newdomainjsonbyte, _ := json.Marshal(newdomain)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.digitalocean.com/v2/domains", bytes.NewBuffer(newdomainjsonbyte))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		fmt.Println("response Status:", response.Status)
		fmt.Println("response Headers:", response.Header)

		retcontents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

	}

	return retcontents
}

func PostNewRecord(token string, newdomain domains.NewDomain, record domains.NewRecord) []byte {

	var retcontents []byte

	newrecordbyte, _ := json.Marshal(record)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.digitalocean.com/v2/domains/"+newdomain.Name+"/records", bytes.NewBuffer(newrecordbyte))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		fmt.Println("response Status:", response.Status)
		fmt.Println("response Headers:", response.Header)

		retcontents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

	}

	return retcontents
}
