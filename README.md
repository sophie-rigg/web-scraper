# web-scraper

This is a simple web scraper takes a domain and lists all possible links within the domain
and all prints all the links on those pages.

## Running the program

To run the program, simply run the following command in the terminal:

```go run cmd/webscraper/main.go --domain=<domain> --level=<level>```

Where `<level>` is the log level and `<domain>` is the domain you want to scrape.

`<level>` can be one of the following: `debug`, `trace`, `info`, `warn`, `error`, `fatal`, `panic`.

The html of the domain must be valid utf8 encoded html.

## Running the tests

To run the tests, simply run the following command in the terminal:

```go test ./...```

The code must also pass the following linters:

```
golangci-lint run --disable-all -E staticcheck -E unused -E errcheck --verbose --timeout 3m
go fmt ./...
go vet ./...
```

## For future development

- Add output not just logs
- Add more tests
- Currently, it only retries a url 3 times, on failure it would stop
- Could be easily blocked by websites, need to add a delay between requests
- Does not handle pop-ups