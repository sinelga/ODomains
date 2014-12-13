package httpgetter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	
)

func GetAll(token string, urlstr string) []byte {

	var retcontents []byte

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlstr, nil)
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		retcontents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} 		

	}

	return retcontents

}
