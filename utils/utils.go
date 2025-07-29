package utils

import "time"

func Sleep(x int) {
	time.Sleep(time.Duration(x) * time.Second)
}
