package main

import "time"

func do(action string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- action
			time.Sleep(1 * time.Second)
		}
	}()
	return c
}

func run(c1, c2 <-chan string) <-chan string {
	outChan := make(chan string)
	go func() {
		for {
			select {
			case r := <-c1:
				outChan <- r
			case r := <-c2:
				outChan <- r
			}
		}
	}()
	return outChan
}

func main() {
	ping := do("ping")
	pong := do("pong")
	game := run(ping, pong)

	for i := 0; i < 10; i++ {
		println(<-game)
	}
}
