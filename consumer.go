// consumer.go
package main

import (
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		attempts := 0
		for {
			err := Sender(recipient)
			if err == nil {
				break
			}
			attempts++
			if attempts >= 3 {
				break
			}
			time.Sleep(time.Duration(1<<attempts) * time.Second)
		}
	}
}
