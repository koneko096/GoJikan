GoJikan
===
[![Travis (.org)](https://img.shields.io/travis/koneko096/gojikan.svg?style=flat-square)](https://travis-ci.org/koneko096/gojikan)
[![Codecov](https://img.shields.io/codecov/c/github/koneko096/gojikan.svg?style=flat-square)](https://codecov.io/gh/koneko096/gojikan/)
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
