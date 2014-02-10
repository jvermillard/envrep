package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	envs := os.Environ()

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	text := string(bytes[:])

	for _, v := range envs {
		t := strings.Split(v, "=")
		text = strings.Replace(text, "[["+t[0]+"]]", t[1], -1)
	}

	fmt.Print(text)
}
