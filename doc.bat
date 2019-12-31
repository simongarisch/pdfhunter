REM https://medium.com/@elliotchance/godoc-tips-tricks-cda6571549b
go get golang.org/x/tools/cmd/godoc
start "" http://localhost:6060/pkg/pdfhunter/
godoc -http=:6060
