package mdtohtml

import (
	"log"
	"os"
)

func LoadHtml(htmlFilePath string) string {
	htmlContent, err := os.ReadFile(htmlFilePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(htmlContent)
}
