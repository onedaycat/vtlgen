build:
	go build -o ./vtlgen/vtlgen ./vtlgen

test:
	go test ./... -v -count=1

bench:
	go test -benchmem -bench=. -v

# correct example
example:
	vtlgen/vtlgen -dir=testdata/mapping-templates -out=testData/mappingTemplates.yml

# manual test
manual:
	vtlgen/vtlgen -dir=testdata/mapping-templates -out=testData/testing.yml

.DEFAULT_GOAL := build