package main

import (
	"compress/gzip"
	"io"
	"log"
	"fmt"
)

func main() {

	r := "<CurrencyConversionRQ Version='1'><InitialValue CurrencyCode='JPY' Amount='1000'/></CurrencyConversionRQ>"

	b := []byte(r)

	var w io.Writer

	gz := gzip.NewWriter(w)
	var i int
	i, err := gz.Write(b)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(i)
	defer gz.Close()
}
