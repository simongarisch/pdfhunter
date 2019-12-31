package pdfhunter

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	LINK = "https://people.gnome.org/~swilmet/glib-gtk-dev-platform.pdf"
)

func TestGetPdfName(t *testing.T) {
	name, err := getPdfName(LINK)

	if err != nil {
		t.Fatalf("error in getPdfName - %s", err)
	}

	expectedName := "glib-gtk-dev-platform.pdf"
	if name != expectedName {
		t.Fatalf("error in name for getPdfName. Expected %q, got %q", expectedName, name)
	}
}

func TestFileExists(t *testing.T) {
	var exists bool

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Cannot get working directory")
	}

	var filePath string
	filePath = "README.md"
	exists, _ = fileExists(filePath)
	if exists == false {
		t.Fatalf("%q file should exist in %q", filePath, dir)
	}

	filePath = "notafile.go"
	exists, _ = fileExists(filePath)
	if exists == true {
		t.Fatalf("%q should not exist in %q", filePath, dir)
	}
}

func TestDownloadFile(t *testing.T) {
	var fileName string
	var exists bool
	var err error

	fileName, err = getPdfName(LINK)
	if err != nil {
		t.Fatalf("error in getPdfName - %s", err)
	}

	exists, err = fileExists(fileName)
	if err != nil {
		t.Fatalf("error in fileExists - %s", err)
	}

	if exists { // file already exists
		err := os.Remove(fileName)
		if err != nil {
			t.Fatalf("error deleting test file - %s", err)
		}
	}

	DownloadFile(fileName, LINK)
	exists, _ = fileExists(fileName)
	if !exists {
		t.Fatal("Downloaded should have been successful")
	}
	os.Remove(fileName)
}

func deleteTestPdfs() error {
	files, err := filepath.Glob(filepath.Join("test_data", "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file, ".pdf") {
			err = os.Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func countPdfs() (int, error) {
	counter := 0

	files, err := filepath.Glob(filepath.Join("test_data", "*"))
	if err != nil {
		return counter, err
	}

	for _, file := range files {
		if strings.HasSuffix(file, ".pdf") {
			if strings.HasSuffix(file, ".pdf") {
				counter++
			}
		}
	}
	return counter, err
}

func serveTestLinks(ready chan bool) {
	http.Handle("/", http.FileServer(http.Dir("./test_data")))
	go http.ListenAndServe(":3001", nil)
	ready <- true
}

func TestDownloadAll(t *testing.T) {
	var count int
	var err error

	err = deleteTestPdfs()
	if err != nil {
		t.Fatalf("error in deleteTestPdfs - %s", err)
	}

	ready := make(chan bool)
	go serveTestLinks(ready)
	<-ready

	err = DownloadAll("test_data", "http://localhost:3001/links.html")
	if err != nil {
		t.Fatalf("error in DownloadAll - %s", err)
	}

	// there are 4 PDF links, so 4 PDF files should be downloaded
	count, err = countPdfs()
	if err != nil {
		t.Fatalf("error in countPdfs - %s", err)
	}

	if count != 4 {
		t.Fatal("expected exactly 4 PDF files")
	}

	err = deleteTestPdfs()
	if err != nil {
		t.Fatalf("error in deleteTestPdfs - %s", err)
	}
}
