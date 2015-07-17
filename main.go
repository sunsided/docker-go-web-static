package main // import "minefield/go-webget"

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://www.google.com")
	check(err)

	body, err := ioutil.ReadAll(response.Body)
	check(err)

	fmt.Printf("%s\n", string(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
