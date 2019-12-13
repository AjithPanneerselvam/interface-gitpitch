package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

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

	return Write(resp.Body, f)
}

func Write(r io.Reader, w io.Writer) error {
	n, err := io.Copy(w, io.TeeReader(r, &WriteCounter{}))
	if err != nil {
		return err
	}

	fmt.Printf("%v written bytes", n)
	return nil
}

func main() {
	err := Download("http://212.183.159.230/1GB.zip", "largeFile.zip")
	if err != nil {
		fmt.Println(err)
	}
}
