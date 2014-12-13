package contentparser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

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
		startpos := strings.Index(zone_file, "IN A")
		endpos := len(zone_file)
		ipstring := zone_file[startpos+5 : endpos]
		domainrecord := []string{domainname, ipstring[:strings.Index(ipstring,"\\n")]}
		domainrecords = append(domainrecords, domainrecord)

	}

	links := jItems["links"].(map[string]interface{})

	pages := links["pages"].(map[string]interface{})

	if next, ok := pages["next"].(string); ok {

		retnext = next

	}

	return retnext,domainrecords 

}
