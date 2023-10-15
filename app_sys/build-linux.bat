set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=1
go build -ldflags "-w -s"  main.go 
pause