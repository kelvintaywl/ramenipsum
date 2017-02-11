# Ramen Ipsum
[![Go Report Card](https://goreportcard.com/badge/github.com/kelvintaywl/ramenipsum)](https://goreportcard.com/report/github.com/kelvintaywl/ramenipsum)

A brothier Lorem Ipsum generator :ramen:

## Motivation

A random thought at work while using the [Bacon Ipsum generator](https://baconipsum.com/) led to this.

A fun project to perhaps commemorate my 3 years in Tokyo, Japan so far.

## Supported Languages

- English :gb: (`en_UK`) (default)
- Japanese :jp: (`ja`)

### Try

> Note: this project uses GoVendor for dependency management.

```
# install govendor
$ go get -u github.com/kardianos/govendor

$ go get github.com/kelvintaywl/ramenipsum
# install dependencies with govendor
$ govendor add +external

$ PORT=9000 go run main.go
```

1. point your browser to http://localhost:9000
2. or grab the random text directly at http://localhost:9000/paragraphs/4 (generates 4 paragraphs of tasty words)
3. grab Japanese text at http://localhost:9000/paragraphs/4?lc=ja
4. grab a bowl of :ramen: and enjoy
