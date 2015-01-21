package networking

import (
	"bufio"
	"fmt"
	"github.com/alexmarchant/torgo/bittorrent"
	"net"
	"time"
)

func SendTrackerRequest(tr *bittorrent.TrackerRequest) error {
	p := make([]byte, 2048)

	fmt.Println("Connecting to: ", tr.URL().String())

	timeout, _ := time.ParseDuration("3s")
	conn, err := net.DialTimeout(tr.URL().Scheme, tr.URL().Host, timeout)
	if err != nil {
		fmt.Println(err)
		return err
	}

	readDeadline := time.Now().Add(timeout)
	err = conn.SetReadDeadline(readDeadline)
	if err != nil {
		return err
	}

	fmt.Println("Sending data")
	fmt.Fprintf(conn, tr.URL().RawQuery)

	_, err = bufio.NewReader(conn).Read(p)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(p)
	fmt.Println(string(p))

	return nil
}
