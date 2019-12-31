// pdfhunter
//
//Download PDF links from a webpage.

package pdfhunter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Example for our package.
func Example() {
	url := "https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books.md"
	folder := "ebooks"
	err := DownloadAll(folder, url)
	if err != nil {
		panic(err)
	}
}

func getPdfName(link string) (string, error) {
	var name string

	if !strings.HasSuffix(link, ".pdf") {
		return name, fmt.Errorf("link %q does point to a pdf file", link)
	}
	parts := strings.Split(link, "/")
	return parts[len(parts)-1], nil
}

func fileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	// has there been an error
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	// there was no error
	if info.IsDir() {
		return false, nil
	}
	return true, nil
}

// DownloadFile downloads a PDF link to a specified file path.
func DownloadFile(filePath string, link string) error {
	exists, _ := fileExists(filePath)
	if exists {
		fmt.Printf("File %q already exists", filePath)
		return nil // don't download the file again
	}

	link = applyURLModifications(link)

	// read link
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	fmt.Println("creating file", filePath)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadAll downloads all PDF links at a particular url address.
func DownloadAll(folder string, url string) error {
	links, err := getLinks(url)
	if err != nil {
		return err
	}

	folderPath := filepath.Join(".", folder)
	os.MkdirAll(folderPath, os.ModePerm) // returns nil if success, else err

	var name string
	for _, link := range links {
		name, err = getPdfName(link)
		if err != nil {
			return err
		}
		filePath := filepath.Join(folderPath, name)
		err := DownloadFile(filePath, link)
		if err != nil {
			fmt.Printf("Unable to download link %q: %q\n", link, err.Error())
		}
	}
	return nil
}
