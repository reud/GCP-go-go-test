package strage

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
)

type CheckList map[string]struct{}

func New(bucketName string, path string) (err error, words string) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return err, ""
	}
	bkt := client.Bucket(bucketName)
	if bkt == nil {
		return errors.New("there is no bucket"), ""
	}

	reader, err := bkt.Object(path).NewReader(context.Background())
	if err != nil {
		return err, ""
	}
	defer deferClose(&err, reader)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		words += scanner.Text() + "\n"
	}
	if err = scanner.Err(); err != nil {
		return err, ""
	}
	return nil, words
}

func deferClose(err *error, fp *storage.Reader) {
	if *err = fp.Close(); *err != nil {
		fmt.Printf("Error:%+v \n", *err)
	}
}
