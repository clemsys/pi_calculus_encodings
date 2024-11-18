// Filename: stats.go
// Description: instrumentation code to gather statistics on  sent/received messages

package stats

import "fmt"

// for instrumentation purposes
type Stats struct {
	sent     int
	received int
	channels int
}

var GlobalStats Stats

func (s *Stats) Reset() {
	s.sent = 0
	s.received = 0
	s.channels = 0
}

func (s *Stats) LogSend() {
	s.sent++
}

func (s *Stats) LogRecv() {
	s.received++
}

func (s *Stats) LogChannel() {
	s.channels++
}

func (s Stats) GetSent() int {
	return s.sent
}

func (s Stats) GetReceived() int {
	return s.received
}

func (s Stats) GetChannels() int {
	return s.channels
}

func (s Stats) PrintStats() {
	fmt.Printf(
		"Stats: %d channels used, %d messages sent, %d messages received\n",
		s.channels,
		s.sent,
		s.received,
	)
}
