# pdfhunter

## Installation

```bash
go get github.com/simongarisch/pdfhunter
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/simongarisch/pdfhunter"
)

func example() {
	url := "https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books.md"
	folder := "ebooks"
	err := pdfhunter.DownloadAll(folder, url)
	if err != nil {
		panic(err)
	}
}

func main() {
	example()
	fmt.Println("All Done...")
}
```