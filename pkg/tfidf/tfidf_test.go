package tfidf

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	os.Setenv("EXECUTION_ENVIRONMENT", "test")
}

func TestCalculateTfIdf(t *testing.T) {
	cmd := CalculateTfIdf()

	b := bytes.NewBufferString("")

	cmd.SetOut(b)

	cmd.SetArgs([]string{"--word", "nunc"})

	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatalf("error in get output: %v", err)
	}

	if out == nil {
		t.Fatalf("error, must be output. expected %s, got %v", "Tf-Idf = [number]", out)
	}
}
