package httpposter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
//	"strings"
	"client"
)

type Command struct {
	
	Data string  `json:"data,omitempty"`
	
}

func PostDomainRecord(token string, newrecord []string) {
	
//	token ="e978588d438fb0b43c9d038ba9f597440b08ed65753610863567137346639085"
	
//	urlstr := "https://api.digitalocean.com/v2/domains/" + newrecord[0] + "/records/1288889"
	urlstr := "https://api.digitalocean.com/v2/domains/erotiikka.org/records/894669"
	
	c,_ := client.New(token)
	
	command := Command{
		
		Data : "104.131.95.7",
		
	}
	
	b, e := json.Marshal(command)
	if e != nil {
		fmt.Printf("%s", e)
	}
	
	fmt.Println(string(b))

//	var jsonStr = bytes.NewBufferString(`{"data":"104.131.95.7"}`)
//	var jsonStr = []byte(`{"data":"104.131.95.7"}`)

//	client := &http.Client{}
	
	req, err := http.NewRequest("POST", urlstr,bytes.NewReader(b))
//	req.Header.Set("Authorization", "Bearer "+token)
//	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	
	fmt.Println("URL ", req.URL.String())
	for k, v := range req.Header {
		fmt.Println("HEADER ", k, v[0])
	}
	
	resp, err := c.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
