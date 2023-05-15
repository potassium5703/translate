// test of github.com/bas24/googletranslate
package main

import (
	"io"
	"os"

	gt "github.com/bas24/googletranslatefree"

	"flag"
	"fmt"
	"strings"
)

func main() {
	english := flag.Bool("en", false, "translate from english")
	var source, into string = "ru", "en"
	var text interface{}
	var err error

	flag.Parse()
	Args := flag.Args()
	if len(Args) > 0 {
		text = strings.Join(Args[:], " ")
	} else {
		text, err = io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		if len(text.([]byte)) == 0 {
			os.Exit(1)
		}
		text = string(text.([]byte)) // looks awful but it works
	}
	if *english {
		source, into = into, source
	}
	result, err := gt.Translate(text.(string), source, into)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
