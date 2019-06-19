package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/koneko096/gojikan"
)

func main() {
	fmt.Print()
	api := gojikan.NewJikan()
	resp, err := api.SearchAnime("Fairy Tail")
	if err != nil {
		log.Fatal(err)
	}

	mam, _ := json.Marshal(resp)
	fmt.Println(string(mam))
}
