package main

import (
	"time"

	"github.com/FuTuL/rmq"
)

func main() {
	connection := rmq.OpenConnection("cleaner", "tcp", "localhost:6379", 2)
	cleaner := rmq.NewCleaner(connection)

	for _ = range time.Tick(time.Second) {
		cleaner.Clean()
	}
}
