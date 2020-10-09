BIN=./exec/cosmic

.PHONY: build_mac
build_mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN) main.go

.PHONY: build_linux
build_linux:
	GOOS=linux GOARCH=amd64 go build -o $(BIN) main.go

.PHONY: clean
clean:
	rm -f $(BIN)