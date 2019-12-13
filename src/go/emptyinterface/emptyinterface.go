package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Downloader ...
type Downloader interface {
	Download(url string) (uuid.UUID, error)
	Pause(downloadID uuid.UUID) error
	Resume(downloadID uuid.UUID) error
	Cancel(downloadID uuid.UUID) error
}

type httpDownloader struct {
	protocol string
}

//type httpDownloader string

func (h *httpDownloader) Download(url string) (uuid.UUID, error) {
	// download ...
	return uuid.New(), nil
}

func (h *httpDownloader) Pause(downloadID uuid.UUID) error {
	// pause the download
	return nil
}

func (h *httpDownloader) Resume(downloadID uuid.UUID) error {
	// resume the download ...
	return nil
}

func (h *httpDownloader) Cancel(downloadID uuid.UUID) error {
	// cancel the download ...
	return nil
}

type torrentDownloader struct {
	protocol string
}

func (h *torrentDownloader) Download(fileID string) (uuid.UUID, error) {
	// download ...
	return uuid.New(), nil
}

func (h *torrentDownloader) Pause(downloadID uuid.UUID) error {
	// pause the download
	return nil
}

func (h *torrentDownloader) Resume(downloadID uuid.UUID) error {
	// resume the download ...
	return nil
}

func (h *torrentDownloader) Cancel(downloadID uuid.UUID) error {
	// cancel the download ...
	return nil
}

func main() {
	var downloader interface{}
	downloader = &httpDownloader{
		protocol: "HTTP/1.1",
	}

	fmt.Println(downloader)
}
