package udp_tracker_protocol

type AnnounceResponse struct {
	Action        actionCode
	TransactionID uint32
	Interval      uint32
	Leechers      uint32
	Seeders       uint32
	IPAddress     uint32
	TCPPort       uint16
}
