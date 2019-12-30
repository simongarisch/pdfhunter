package pdfhunter

import (
	"os"
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

func TestDownloadAll(t *testing.T) {
}
