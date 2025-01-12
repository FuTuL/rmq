package main

import (
	"github.com/FuTuL/rmq"
)

func main() {
	connection := rmq.OpenConnection("cleaner", "tcp", "localhost:6379", 2)
	queue := connection.OpenQueue("things")
	queue.PurgeReady()
}
