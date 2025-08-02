package main

import (
	"github.com/prithivilaksh/transactions/channels"
	"github.com/prithivilaksh/transactions/mutex"
)

func main() {
	channels.SimulateChannels(1000, 1000000)
	mutex.SimulateMutex(1000, 1000000)
}	
