package main

 import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

 const (
	chunkSize = 100000 // Number of lines to process per chunk
)

func main() {
	filePath := "path/to/your/csv/file.csv" // Replace with your file path
 	for {
		err := processCSVChunks(filePath)
		if err != nil {
			log.Printf("Error processing CSV: %s", err.Error())
		}
 		time.Sleep(30 * time.Minute) // Wait for 30 minutes before processing again
	}
}

func processCSVChunks(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %s", err.Error())
	}
	defer file.Close()
 	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record
 	for {
		chunk, err := readCSVChunk(reader)
		if err != nil {
			return fmt.Errorf("failed to read CSV chunk: %s", err.Error())
		}
 		if len(chunk) == 0 {
			break // Reached end of file
		}
 		// Process the chunk and extract the data
		for _, record := range chunk {
			// Process each record as needed
			// Example: fmt.Println(record)
		}
	}
 	return nil
}

func readCSVChunk(reader *csv.Reader) ([][]string, error) {
	chunk := make([][]string, 0, chunkSize)
 	for i := 0; i < chunkSize; i++ {
		record, err := reader.Read()
		if err != nil {
			return nil, err
		}
 		chunk = append(chunk, record)
	}
 	return chunk, nil
}