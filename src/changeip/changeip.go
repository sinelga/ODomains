package changeip

import (
	"encoding/json"
	"fmt"
	"httpgetter"
//	"httpposter"
	"os"
	"strconv"
//	"dumpintocsv"
)

func ChangeFromTo(token string, domain string, fromip string, toip string) []string {

	urlstr := "https://api.digitalocean.com/v2/domains/" + domain + "/records"
	contents := httpgetter.GetAll(token, urlstr)
//	fmt.Printf("%s\n", string(contents))

	var jItems map[string]interface{}

	err := json.Unmarshal(contents, &jItems)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	
//	var recordstodump [][]string
	var newrecord []string
	
	items := jItems["domain_records"].([]interface{})
	
	for _, item := range items {
		
		record := item.(map[string]interface{})
		
		record_type := record["type"].(string)
		 
		 if record_type == "A" {
		 	
		 	record_id :=record["id"].(float64)
		 	
		 	record_id_str :=strconv.FormatFloat(record_id, 'f', 0, 64)
		 	
		 	newrecord =[]string{domain,record_id_str,toip,}
		 	
//		 	recordstodump =append(recordstodump,newrecord) 
		 	
//		 	httpposter.PostDomainRecord(token,newrecord) 
		 	
		 }
		 		
	}
	
//	dumpintocsv.Dump(recordstodump)
	
	return newrecord

}
