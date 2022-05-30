package downloaddocument

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestDownloadDocument(t *testing.T) {

	cmd := DownloadDocument()

	b := bytes.NewBufferString("")

	cmd.SetOut(b)

	cmd.SetArgs([]string{"--url", "https://www.caelum.com.br/apostila/apostila-html-css-javascript.pdf"})

	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatalf("error in get output: %v", err)
	}

	if out == nil {
		t.Fatalf("error, must be output. expected %s, got %v", "Downloaded a file apostila-html-css-javascript.pdf with size 9247560[]", out)
	}
}
