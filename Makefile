build:
	go build -o ./vtlgen/vtlgen ./vtlgen

test:
	go test ./... -v -count=1

.DEFAULT_GOAL := build