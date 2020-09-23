package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://gutenberg.org/cache/epub/16328/pg16328.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("mybook.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	fmt.Println("Text copied to file check it ", file.Name())

}
