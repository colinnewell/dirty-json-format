package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/colinnewell/json-format/json"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	json.Indent(&b, in, "", "  ")
	fmt.Println(b.String())
}
