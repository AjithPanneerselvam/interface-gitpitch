package main

import (
	"net/url"

	"github.com/google/uuid"
)

// Downloader ...
type Downloader interface {
	Download(fileURL *url.URL) (uuid.UUID, error)
	Pause(downloadID uuid.UUID) error
	Resume(downloadID uuid.UUID) error
	Cancel(downloadID uuid.UUID) error
}
