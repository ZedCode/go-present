# go-present Makefile
#
# This file will make go-present even if
# you don't have Go installed on your
# local system.

linux: main.go
	docker run --rm -v $(shell pwd):/usr/src/go-present -w /usr/src/go-present golang:1.8 go build -v
	# We should also rebuild the main container.
	docker build -t="go-present:latest" .
