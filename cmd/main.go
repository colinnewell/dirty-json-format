package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/colinnewell/json-format/json"
)

func main() {
	var b bytes.Buffer
	err := json.Indent(&b, []byte(`{"blah":3`), "", "  ")
	fmt.Println(b.String())
	if err != nil {
		log.Fatal(err)
	}
}
