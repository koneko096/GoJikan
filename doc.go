/*
A simple Go wrapper for unofficial jikan.moe API

Installation

	go get -u github.com/koneko096/gojikan

Usage

	import "github.com/koneko096/gojikan"

	func main() {
	  client := gojikan.NewJikan()
	  anime, _ := client.GetAnime(123)
	  // ...
	}

License

GPLv3

*/
package gojikan
