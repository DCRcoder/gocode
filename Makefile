build:
	go build -o ./bin/gocode -v ./cmd/gocode.go 
install:
	go install ./cmd/gocode.go
