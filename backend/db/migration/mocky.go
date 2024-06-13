package migration

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type Mocky interface {
	// This is an interface for the mocky-cli migration.
	FetchMigrationRecords() []string
	SaveMigrationRecords([]string)
}

type mocky struct {
	// This is a struct for the mocky-cli migration.
}

func NewMocky() Mocky {
	// This is a constructor for the mocky-cli migration.
	return &mocky{}
}

func (m mocky) Run() {
	// Get the migration records from the S3 bucket.
	var files []string = []string{"products-1.jsonl", "products-2.jsonl", "products-3.jsonl", "products-4.jsonl"}

	out := make(chan []models.Product, len(files))

	for _, file := range files {
		go func(file string) {
			resp, err := m.S3Service.GetObjectRequest(os.Getenv("AWS_BUCKET_NAME"), file)
			if err != nil {
				return
			}

			defer resp.Body.Close()
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(resp.Body); err != nil {
				return
			}

			out <- parseMigrationRecords(buf.String())
			time.Sleep(10 * time.Second)
		}(file)
	}

	for i := 0; i < len(files); i++ {
		data := <-out
		// split the data into chunks of 65535
		batchSize := 5000
		for i := 0; i < len(data); i += batchSize {
			end := i + batchSize

			if end > len(data) {
				end = len(data)
			}

			m.insertProducts(data[i:end])

			fmt.Println("Inserted", end-i, "products")
		}

	}

}
