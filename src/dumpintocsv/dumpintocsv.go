package dumpintocsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Dump(recordtomodify [][]string) {

	file, error := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, 0666)
	if error != nil {
		panic(error)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, record := range recordtomodify {

		fmt.Println(record[0])
		fmt.Println(record[1])
		fmt.Println(record[2])


	}

	returnError := writer.WriteAll(recordtomodify)
	if returnError != nil {
		fmt.Println(returnError)
	} else {

		writer.Flush()

	}

}
