package pdfhunter


func getLinks(url string) ([]string, error) {
	links := []string{}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return links, err
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		if strings.HasSuffix(href, ".pdf") {
			links = append(links, href)
		}
	})
	return links, nil
}