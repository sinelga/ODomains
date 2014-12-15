package contentparser

import (
	"encoding/json"
	"fmt"
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



func Parse(contents []byte) (string, [][]string){

	var retnext string

	var jItems map[string]interface{}

	err := json.Unmarshal(contents, &jItems)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	items := jItems["domains"].([]interface{})

	var domainrecords [][]string

	for _, ii := range items {

		iii := ii.(map[string]interface{})
		domainname := iii["name"].(string)

		zone_file := iii["zone_file"].(string)
//		startpos := strings.Index(zone_file, "IN A")
//		endpos := len(zone_file)
//		ipstring := zone_file[startpos+5 : endpos]
//		fmt.Println(zone_file)
		ipstring :=findIP(strings.TrimSpace(zone_file))
		
		fmt.Println(ipstring)
		
		if domainname == "k18.me" {
			
			fmt.Println(zone_file)
			
		}
				
		
//		domainrecord := []string{domainname, ipstring[:strings.Index(ipstring,"\\n")]}
		domainrecord := []string{domainname,strings.TrimSpace(ipstring)} 
		domainrecords = append(domainrecords, domainrecord)

	}
	
	fmt.Println("all domains ",len(domainrecords))

	links := jItems["links"].(map[string]interface{})

	pages := links["pages"].(map[string]interface{})

	if next, ok := pages["next"].(string); ok {

		retnext = next

	}

	return retnext,domainrecords 

}
