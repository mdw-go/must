#!/usr/bin/make -f

test:
	go fmt ./...
	go mod tidy
	go test -cover -timeout=1s -race ./...

onefile:
	@go-mergepkg -dirs "." -header "github.com/mdwhatcott/must@$(shell git describe) (a little copy-paste is better than a little dependency)"
