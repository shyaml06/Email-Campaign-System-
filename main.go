// main.go
package main

import (
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go func() {
		loadRecipient("./email.csv", recipientChannel)
	}()

	var wg sync.WaitGroup
	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()
}
