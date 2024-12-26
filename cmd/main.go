package main

import (
	"own-redis/internal/eventloop"
)

func main() {
	loop := eventloop.NewEventLoop()
	loop.Start("6379")
}
