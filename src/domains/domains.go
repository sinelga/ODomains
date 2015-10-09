package domains

import ()

type Domaincsv struct {
	Locale string
	Themes string
	Name   string
	Ip     string
}
type NewDomain struct {
	Name string `json:"name"`
	Ip   string `json:"ip_address"`
}

type NewRecord struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority string `json:"priority"`
	Port     string `json:"port"`
	Weight   string `json:"weight"`
}
