go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o .\releases\nekosrun-windows-amd64.exe
go env -w GOARCH=i386
go build -o .\releases\nekosrun-windows-i386.exe
go env -w GOARCH=arm64
go build -o .\releases\nekosrun-windows-arm64.exe

go env -w GOOS=linux
go env -w GOARCH=amd64
go build -o .\releases\nekosrun-linux-amd64
go env -w GOARCH=i386
go build -o .\releases\nekosrun-linux-i386
go env -w GOARCH=arm
go build -o .\releases\nekosrun-linux-mips
go env -w GOARCH=arm64
go build -o .\releases\nekosrun-linux-arm
go env -w GOARCH=mips
go build -o .\releases\nekosrun-linux-mips

go env -w GOOS=darwin
go env -w GOARCH=amd64
go build -o .\releases\nekosrun-darwin-amd64
go env -w GOARCH=amd64
go build -o .\releases\nekosrun-darwin-arm64
