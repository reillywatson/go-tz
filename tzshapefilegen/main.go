package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file := flag.String("file", "", "file to embed")
	name := flag.String("name", "tzShapeFile", "var name")
	flag.Parse()
	if *file == "" {
		log.Fatalln("No file given")
	}
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("could not read file: %v\n", err)
	}
	defer f.Close()
	buf := bytes.NewBuffer([]byte{})
	g, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		log.Printf("could not create gzip writer: %v\n", err)
		return
	}
	defer g.Close()
	_, err = io.Copy(g, f)
	if err != nil {
		log.Printf("could not copy data: %v\n", err)
		return
	}
	if err := g.Flush(); err != nil {
		log.Printf("could not flush gzip: %v\n", err)
		return
	}
	str := buf.Bytes()
	out := bytes.NewBuffer([]byte{})
	for i := range str {
		if int(str[i]) < 16 {
			out.WriteString("\\x" + fmt.Sprintf("0%X", str[i]))
		} else {
			out.WriteString("\\x" + fmt.Sprintf("%X", str[i]))
		}
	}
	var template = `package gotz

var %s = []byte("%s")
`
	content := fmt.Sprintf(template, *name, out)
	fout, err := os.Create("tzshapefile.go")
	if err != nil {
		log.Printf("could not create tzshapefile.go: %v", err)
		return
	}
	defer fout.Close()
	_, err = fout.WriteString(content)
	if err != nil {
		log.Printf("could not write content: %v", err)
		return
	}
}
