package downloader

import (
	"fmt"
	"github.com/alexmarchant/torgo/bittorrent"
	"github.com/alexmarchant/torgo/networking"
)

var (
	Port   = 6881
	PeerID = "15620985492012023883"
)

type Downloader struct {
	Torrent      *bittorrent.Torrent
	DownloadPath string
}

func NewDownloader(torrentPath string, downloadPath string) (downloader *Downloader, err error) {
	var torrent *bittorrent.Torrent

	torrent, err = bittorrent.NewTorrent(torrentPath)
	if err != nil {
		return
	}

	downloader = &Downloader{
		Torrent:      torrent,
		DownloadPath: downloadPath,
	}

	return
}

func (d *Downloader) StartDownload() (err error) {
	for _, tracker := range d.Torrent.Trackers() {
		trackerRequest := tracker.Request(d.Torrent.InfoHash(), PeerID, Port)
		tErr := networking.SendConnectRequest(trackerRequest)
		if tErr != nil {
			fmt.Println(tErr)
		}
	}
	return
}
