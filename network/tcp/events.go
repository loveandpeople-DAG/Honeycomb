package tcp

import "github.com/loveandpeople-DAG/goHive/events"

type tcpServerEvents struct {
	Start    *events.Event
	Shutdown *events.Event
	Connect  *events.Event
	Error    *events.Event
}
