package main

import "sync"

func fixed() {
	// START OMIT
	wg := sync.WaitGroup{}
	wg.Add(3)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // HLcopy

		go func() {
			println(v)
			wg.Done()
		}()
	}
	
	wg.Wait()
	// END OMIT
}
