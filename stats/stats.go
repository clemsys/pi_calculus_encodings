package stats

import "fmt"

// for instrumentation purposes
type Stats struct {
	sent     int
	received int
}

var GlobalStats Stats

func (s *Stats) Reset() {
	s.sent = 0
	s.received = 0
}

func (s *Stats) LogSend() {
	s.sent++
}

func (s *Stats) LogRecv() {
	s.received++
}

func (s Stats) GetSent() int {
	return s.sent
}

func (s Stats) GetReceived() int {
	return s.received
}

func (s Stats) PrintStats() {
	fmt.Printf("Stats: %d messages sent, %d messages received\n", s.sent, s.received)
}
