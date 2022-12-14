package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	p := fmt.Println
	file, _ := os.Create("file.txt")
	file.WriteString("hello therer...")
	file.Close()
	//read the file
	stream, _ := ioutil.ReadFile("file.txt")
	fmt.Println(stream)
	a := string(stream)
	p(a)
	b := []byte(a)
	p(b)

	if "japan" == "JAPAN" {
		p(true)
	} else {
		p(false)
	}
	aa := strings.EqualFold("japan", "JAPAN")
	p(aa)
	for i, ch := range "apple" {
		p("%d : %s \n", i, ch)
	}
}
