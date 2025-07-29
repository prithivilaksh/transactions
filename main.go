package main

import (
	"github.com/prithivilaksh/transactions/channels"
	"github.com/prithivilaksh/transactions/mutex"
)

func main() {
	mutex.SimulateMutex()
	channels.SimulateChannels()
}
