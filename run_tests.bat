go get golang.org/x/tools/cmd/cover
go test -coverprofile cover.out
go tool cover -html=cover.out
pause