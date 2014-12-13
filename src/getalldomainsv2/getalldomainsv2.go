package getalldomainsv2

import (
	"contentparser"
	"httpgetter"
)

func GetAllD(token string, urlstr string) [][]string {

	var alldomainrecords [][]string

	contents := httpgetter.GetAll(token, urlstr)

	nextlink, domainsrecords := contentparser.Parse(contents)

	for _, record := range domainsrecords {

		alldomainrecords = append(alldomainrecords, record)

	}

	for {

		if nextlink != "" {
			contents = httpgetter.GetAll(token, nextlink)
			nextlink, domainsrecords = contentparser.Parse(contents)

			for _, record := range domainsrecords {

				alldomainrecords = append(alldomainrecords, record)

			}
		} else {
			{
				break
			}
		}

	}

	return alldomainrecords

}
