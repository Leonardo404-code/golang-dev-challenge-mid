# CLI TFIDF and File Downloader

Test done successfully!
This is my first CLI created, I hope to practice on other projects and improve on this type of project!

## About architecture

The architecture I usually use is based on the [project-layout](https://github.com/golang-standards/project-layout) repository from [golang-standards](https://github.com/golang-standards).

I will describe details of the functions of each folder

### cmd

contains the executable code of the application, only the main one that starts the application will remain in it

### config

Defines variables, constants and functions that will change the application configuration

### data

contains files, pdf, txt and etc

### docs

Design and user documents (in addition to your godoc generated documentation).

### pkg

Library code that's ok to use by external applications or by the application itself for helper codes

## How works

clone repository

```bash
git clone https://github.com/Leonardo404-code/golang-dev-challenge-mid.git
```

install all dependencies

```go
go mod tidy
```

Build application

```go
go build ./cmd/tfidf/
```

Run application

```bash
./tfidf
```

And see the help text

```bash
Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  down        download file from url
  help        Help about any command
  tfidf       return TFIDF of word

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

We see the two program commands, ifidf and down

## TFIDF

the tfidf command receives a flag called word which will be the word analyzed in the three documents in the data folder. see a running example

```bash
./tfidf tfidf --word nunc
```

expected return

```bash
Tf-Idf = -0.000375956707333744
```

See help command for TFIDF

```bash
./tfidf tfidf -h
```

And see more information about its feature

```bash
return TFIDF of word

Usage:
   tfidf --word [word] [flags]

Examples:
./tfidf tfidf --word hello

Flags:
  -h, --help          help for tfidf
  -w, --word string   flag for word information
```

## Download Files

the down command receives a flag named url, this flag must point to a valid domain and will download the found document. see a running example

```bash
./tfidf down --url https://www.golang-book.com/public/pdf/gobook.pdf
```

expected return

```bash
Downloaded a file gobook.pdf with size 2904832%
```

The file appears in the root of your project

See help command for Download Files

```bash
./tfidf down -h
```

And see more information about its feature

```bash
download file from url

Usage:
   down --url [url] [flags]

Examples:
./tfidf down --url https://www.golang-book.com/public/pdf/gobook.pdf

Flags:
  -h, --help         help for down
  -u, --url string   document url flag (default "https://www.golang-book.com/public/pdf/gobook.pdf")
```

## Unit tests

Unit tests were performed for each feature of the application. to start them type in root folder

```go
go test ./...
```
