package main

import (
	"sync"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		Sender(recipient)

	}

}
