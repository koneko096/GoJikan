GoJikan
===
![License](https://img.shields.io/github/license/koneko096/gojikan.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/koneko096/gojikan?status.svg)](https://godoc.org/github.com/koneko096/gojikan)
[![Travis (.org)](https://img.shields.io/travis/koneko096/gojikan.svg?style=flat-square)](https://travis-ci.org/koneko096/gojikan)
[![Codecov](https://img.shields.io/codecov/c/github/koneko096/gojikan.svg?style=flat-square)](https://codecov.io/gh/koneko096/gojikan/)
[![GoReport](https://goreportcard.com/badge/github.com/koneko096/gojikan)](https://goreportcard.com/report/github.com/koneko096/gojikan)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

A simple Go wrapper for unofficial jikan.moe API

## Installation
```shell
$ go get -u github.com/koneko096/gojikan
```

## Usage
```go
import "github.com/koneko096/gojikan"

func main() {
  client := gojikan.NewJikan()
  anime := client.GetAnime(123)
  ...
}
```
