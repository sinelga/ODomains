package getalldomains

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"regexp"
)

func findIP(input string) string {
         numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
         regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

         regEx := regexp.MustCompile(regexPattern)
         return regEx.FindString(input)
 }


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
				
				ipstring :=findIP(strings.TrimSpace(jlive_zone_file))
				
//				fmt.Println(domainname)
//				fmt.Println(iii)
//				fmt.Println(jlive_zone_file)

				fmt.Println(domainname,	ipstring)			
				 							
			}
			
			
				
//		jItems["status"]["domains"].([]interface{})	

//		itemsii := items["domains"].([]interface{})


//		items := jItems["status"]["domains"]
		
//		fmt.Println(items) 

	}

}
