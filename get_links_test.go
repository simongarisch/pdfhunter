package pdfhunter

import "testing"

const (
	URL = "https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books.md"
)

func TestSearch(t *testing.T) {
	links, err := getLinks(URL)

	if err != nil {
		t.Fatalf("error in getLinks - %s", err)
	}

	if len(links) == 0 {
		t.Fatal("There should be at least one pdf link")
	}
}
