package main

import (
	downloaddocument "golang-challenge/pkg/download_document"
	"golang-challenge/pkg/tfidf"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}

	cmd.AddCommand(tfidf.CalculateTfIdf())

	cmd.AddCommand(downloaddocument.DownloadDocument())

	cmd.Execute()
}
