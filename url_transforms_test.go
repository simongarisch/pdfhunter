package pdfhunter

import "testing"

func TestApplyURLModificationsGithub(t *testing.T) {
	var url, expectedResult string
	url = applyURLModifications("https://github.com/.../blob/master/....pdf")
	expectedResult = "https://github.com/.../raw/master/....pdf"
	if url != expectedResult {
		t.Fatalf("expected %q, got %q", expectedResult, url)
	}

	url = applyURLModifications("https://this_will_not_change.pdf")
	expectedResult = "https://this_will_not_change.pdf"
	if url != expectedResult {
		t.Fatalf("expected %q, got %q", expectedResult, url)
	}
}
