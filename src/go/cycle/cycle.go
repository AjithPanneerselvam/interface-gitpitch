package main

type Downloader interface {
	Download()
}

type TorrentDownloader interface {
	Downloader
}
