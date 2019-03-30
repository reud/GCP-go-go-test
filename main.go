package main

import (
	"fmt"
	"github.com/reud/GCP-go-go/strage"
	"os"
)

func main() {
	err, words := strage.New(os.Getenv("BUCKET_NAME"), os.Getenv("FILE_PATH"))
	if err != nil {
		panic(err)
	}
	fmt.Print(words)
}
