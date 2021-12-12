# build linux amd64 executable
$env:GOOS="linux"
$env:GOARCH="amd64"

go build -o address-server-linux-amd64

# build windows amd64 executable
$env:GOOS="windows"
$env:GOARCH="amd64"

go build -o address-server-windows-amd64.exe
