package main

import "time"

var Now = func() time.Time {
	return time.Now()
}
