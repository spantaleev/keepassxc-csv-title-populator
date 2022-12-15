package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Expected 2 arguments: SOURCE_FILE.csv DEST_FILE.csv")
		os.Exit(1)
	}

	srcFilePath := os.Args[1]
	dstFilePath := os.Args[2]

	inputRecords, err := readCsvFromFilePath(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}

	titleColumnIndex := getColumnIndexByValue(inputRecords[0], "Title")
	if titleColumnIndex == -1 {
		log.Fatal("Could not find a `Title` column")
	}

	urlColumnIndex := getColumnIndexByValue(inputRecords[0], "URL")
	if titleColumnIndex == -1 {
		log.Fatal("Could not find a `URL` column")
	}

	outputRecords := make([][]string, len(inputRecords))

	for idx, record := range inputRecords {
		if idx == 0 {
			outputRecords[idx] = record
			continue
		}

		if record[titleColumnIndex] != "" {
			outputRecords[idx] = record
			continue
		}

		// No title. Let's come up with something

		record[titleColumnIndex] = determineTitleByURL(record[urlColumnIndex])

		outputRecords[idx] = record
	}

	writeRecordsToCSVFile(outputRecords, dstFilePath)
	// fmt.Println(titleColumnIndex)
	// fmt.Println(urlColumnIndex)
	// fmt.Println(outputRecords)

	// return records

}

func readCsvFromFilePath(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to parse file as CSV for "+filePath, err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("no records found - nothing to do")
	}

	return records, nil
}

func writeRecordsToCSVFile(records [][]string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file "+filePath, err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	return w.WriteAll(records)
}

func getColumnIndexByValue(columns []string, match string) int {
	for idx, value := range columns {
		if value == match {
			return idx
		}
	}
	return -1
}

func determineTitleByURL(urlString string) string {
	parsed, err := url.Parse(urlString)
	if err != nil {
		return "Unknown"
	}

	return parsed.Hostname()
}
