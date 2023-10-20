// test of github.com/bas24/googletranslate
package main

import (
	"bufio"
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
	var text string
	var err error

	flag.Parse()
	if *english {
		source, into = into, source
	}

	Args := flag.Args()
	if len(Args) > 0 {
		text = strings.Join(Args[:], " ")
	} else {
		r := bufio.NewReader(os.Stdin)
		for {
			text, err = r.ReadString('\n')
			if err != nil {
				switch err {
				case io.EOF:
					return
				default:
					panic(err)
				}
			}
			text = strings.Trim(text, "\n")
			if len(text) == 0 {
				continue
			}
			result, err := gt.Translate(text, source, into)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
		}
	}
	result, err := gt.Translate(text, source, into)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
