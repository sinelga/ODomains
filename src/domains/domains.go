package domains

import (

)

type Domaincsv struct {
	Locale string
	Themes string
	Name string
	Ip string
	
}
type NewDomain struct {
	
	Name string `json:"name"`
	Ip string   `json:"ip_address"`
	
}