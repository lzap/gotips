package main

import "time"

func main() {
	for i := range 3 {
		go func() {
			println(i)
		}()
	}

	time.Sleep(100 * time.Millisecond)
}
