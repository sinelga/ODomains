package httpposter

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

var token string

func TestPostDomainRecord(t *testing.T) {

	csvFile, err := os.Open("/home/juno/git/OceanDomains/OceanDomains/token.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(csvFile)

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		token = fields[0]

	}

	newrecord := []string{"alastomia.com", "1288889", "104.131.95.7"}

	PostDomainRecord(token, newrecord)

}
