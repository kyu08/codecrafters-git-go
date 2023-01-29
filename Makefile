.PHONY: test

test:
		go test -v ./... | cgt
