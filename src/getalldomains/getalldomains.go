package getalldomains

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetAllD(your_client_id string, api_key string) {

	urlstr := "https://api.digitalocean.com/domains?client_id=" + your_client_id + "&api_key=" + api_key
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
//		fmt.Printf("%s\n", string(contents))

		var jItems map[string]interface{}

		err = json.Unmarshal(contents, &jItems)
		if err != nil {
			//			golog.Err(err.Error())
			fmt.Printf("%s", err)
			os.Exit(1)
		}

			items := jItems["domains"].([]interface{})
			
			for _, ii := range items {
			
				iii := ii.(map[string]interface{})
				domainname := iii["name"].(string)
//				fmt.Println(domainname)
				
				jlive_zone_file := iii["live_zone_file"].(string)
				
				pos :=strings.Index(jlive_zone_file,"IN A") 
				fmt.Println(domainname,jlive_zone_file[pos:pos+20])				
				 
				
			
			}
			
			
				
//		jItems["status"]["domains"].([]interface{})	

//		itemsii := items["domains"].([]interface{})


//		items := jItems["status"]["domains"]
		
//		fmt.Println(items) 

	}

}
