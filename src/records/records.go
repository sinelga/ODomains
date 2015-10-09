package records

import (
//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

)

func AddCname(your_client_id string, api_key string, domainid int,name string, hostname string) {

	urlstr := "https://api.digitalocean.com/domains/" +  strconv.Itoa(domainid)+ "/records/new?client_id=" + your_client_id + "&api_key=" + api_key + "&record_type=CNAME"  + "&name=" +name+"&data="+hostname
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