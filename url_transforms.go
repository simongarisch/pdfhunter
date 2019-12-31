package pdfhunter

import "strings"

func modifyGithubURL(url string) string {
	if strings.HasPrefix(url, "https://github.com/") {
		return strings.ReplaceAll(url, "/blob/", "/raw/")
	}
	return url
}

func applyURLModifications(url string) string {
	allModifications := []func(string) string{
		modifyGithubURL,
	}
	for _, modfunc := range allModifications {
		url = modfunc(url)
	}
	return url
}
