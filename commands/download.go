package commands

import (
	"github.com/alexmarchant/torgo/downloader"
	"os/user"
)

func Download(cmd *Command) error {
	if len(cmd.Options) < 1 {
		return BadCommand
	}

	torrentPath := cmd.Options[0]
	downloadPath := downloadPath(cmd)

	myDownloader, err := downloader.NewDownloader(torrentPath, downloadPath)
	if err != nil {
		return err
	}

	err = myDownloader.StartDownload()
	if err != nil {
		return err
	}

	return nil
}

func downloadPath(cmd *Command) string {
	if len(cmd.Options) < 2 {
		return defaultDownloadPath()
	} else {
		return cmd.Options[1]
	}
}

func defaultDownloadPath() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/Downloads"
}
