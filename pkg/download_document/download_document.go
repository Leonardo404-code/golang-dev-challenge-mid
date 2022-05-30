package downloaddocument

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DownloadDocument find url of document in web and download the content
func DownloadDocument() *cobra.Command {
	var (
		urlPath, fileName string
	)

	cmd := &cobra.Command{
		Use:     "down --url [url]",
		Short:   "download file from url",
		Example: "./tfidf down --url https://www.golang-book.com/public/pdf/gobook.pdf",
		Run: func(cmd *cobra.Command, args []string) {
			fileUrl, err := url.Parse(urlPath)

			if err != nil {
				log.Fatalf("error in parse url: %v", err)
			}

			path := fileUrl.Path

			if path == "/" {
				log.Fatal("URL is not from a PDF")
			}

			segments := strings.Split(path, "/")

			fileName = segments[len(segments)-1]

			file, err := os.Create(fileName)

			if err != nil {
				log.Fatalf("error in create file: %v", err)
			}

			resp, err := http.Get(urlPath)

			if err != nil {
				log.Fatalf("error in request: %v", err)
			}

			defer resp.Body.Close()

			size, err := io.Copy(file, resp.Body)

			if err != nil {
				log.Printf("error in calculate size: %v", err)
			}

			defer file.Close()

			fmt.Printf("Downloaded a file %s with size %d", fileName, size)
		},
	}

	cmd.Flags().StringVarP(&urlPath, "url", "u", "https://www.golang-book.com/public/pdf/gobook.pdf", "document url flag")

	return cmd
}
