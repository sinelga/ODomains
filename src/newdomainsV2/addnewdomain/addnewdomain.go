package addnewdomain

import (

"domains"
//"fmt"
"httpgetter"
//"encoding/json"
)

func Add(token string,domaincsv domains.Domaincsv) {
	
//	fmt.Println(domaincsv)
	
	newdomain :=domains.NewDomain{
		
		Name: domaincsv.Name,
		Ip: domaincsv.Ip,  
		
	}
	
	
	httpgetter.PostNewDomain(token,newdomain)		 
	
}