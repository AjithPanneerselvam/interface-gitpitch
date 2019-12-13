package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url string, downloadPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(downloadPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	return Write(resp, f)
}

func Write(resp *http.Response, f *os.File) error {
	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%v written bytes", n)
	return nil
}

func main() {
	err := Download("https://tinyurl.com/sxynmjp", "./goTenthAnniversary.png")
	if err != nil {
		fmt.Println(err)
	}
}
