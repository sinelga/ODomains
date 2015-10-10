package add_record_infile

import (
    "testing"
    "domains"
)

func TestAdd(t *testing.T) {
	
	domaincsv :=domains.Domaincsv{
		
		Locale: "fi_FI",
		Themes: "porno",
		Name: "test.com",
		Ip: "1.1.1.1",
		
	}
	
	Add(domaincsv)
	

}

