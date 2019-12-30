package pdfhunter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

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

func DownloadFile(filePath string, link string) error {
	exists, _ := fileExists(filePath)
	if exists {
		fmt.Printf("File %q already exists", filePath)
		return nil // don't download the file again
	}

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