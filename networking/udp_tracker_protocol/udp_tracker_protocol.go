package udp_tracker_protocol

import (
	"time"
)

var (
	timeout, _               = time.ParseDuration("3s")
	blankConnectionID uint64 = 4497486125440
)

type actionCode uint32

const (
	connectAction actionCode = iota
	announceAction
	scrapeAction
	errorAction
)

type eventCode uint32

const (
	noneEvent eventCode = iota
	completedEvent
	startedEvent
	stoppedEvent
)
