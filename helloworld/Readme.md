sudo apt install -y golang
go mod init example/hello -> creates go.mod
create hello.go and add external packages
go mod tidy
go run .
go mod edit -replace example.com/greetings=../greetings // The command specifies that example.com/greetings should be replaced with ../greetings for the purpose of locating the dependency
