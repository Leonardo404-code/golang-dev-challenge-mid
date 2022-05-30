package tfidf

import (
	"fmt"
	"golang-challenge/config"
	"io/fs"
	"log"
	"math"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CalculateTfIdf return TF and IDF of a word
func CalculateTfIdf() *cobra.Command {
	var (
		tf, idf, countString, contentLarge, totalWordInDocument float64
		word                                                    string
	)

	cmd := &cobra.Command{
		Use:     "tfidf --word [word]",
		Short:   "return TFIDF of word",
		Example: "./tfidf tfidf --word hello",
		Run: func(cmd *cobra.Command, args []string) {
			if word == "" {
				log.Fatal("--word flag can not be empty")
			}

			var (
				isTest bool
				files  []fs.DirEntry
				err    error
			)

			isTest = config.IsTestRun()

			if isTest {
				files, err = os.ReadDir("../../data")

				if err != nil {
					log.Fatalf("error in load dir: %v", err)
				}
			} else {
				files, err = os.ReadDir("./data")

				if err != nil {
					log.Fatalf("error in load dir: %v", err)
				}
			}

			totalDocuments := float64(len(files))

			for _, file := range files {
				var content []byte

				if isTest {
					content, err = os.ReadFile(fmt.Sprintf("../../data/%s", file.Name()))

					if err != nil {
						log.Fatalf("error in read file %v", err)
					}
				} else {
					content, err = os.ReadFile(fmt.Sprintf("./data/%s", file.Name()))

					if err != nil {
						log.Fatalf("error in read file %v", err)
					}
				}

				convertContent := string(content)

				countString += float64(strings.Count(strings.ToLower(convertContent), strings.ToLower(word)))

				contentLarge += float64(len(convertContent))

				if strings.Count(strings.ToLower(convertContent), strings.ToLower(word)) > 0 {
					totalWordInDocument += 1
				}
			}

			tf = countString / contentLarge

			idf = math.Log(totalDocuments / (totalWordInDocument + 1))

			tfIdf := tf * idf

			log.Printf("Tf-Idf = %v", tfIdf)
		},
	}

	cmd.Flags().StringVarP(&word, "word", "w", "", "flag for word information")

	return cmd
}
