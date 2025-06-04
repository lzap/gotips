package main

func f(i int) {
	if i--; i <= 0 {
		return
	}
	f(i)
}
