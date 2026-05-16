package main

import (
	"fmt"
	"sync"
	"time"
)

	type Broker struct {
		subChan   chan chan int
		unsubChan chan chan int
		dataChan  chan int
	}

	func Init() *Broker {
		b := &Broker{
			subChan:   make(chan chan int),
			unsubChan: make(chan chan int),
			dataChan:  make(chan int),
		}

		go b.Serve()

		return b
	}

	func (b *Broker) Serve() {
		subs := make(map[chan int]struct{})
		for {
			select {
			case ch := <-b.subChan:
				subs[ch] = struct{}{}
			case ch := <-b.unsubChan:
				delete(subs, ch)
			case data, ok := <-b.dataChan:
				if !ok {
					for sub := range subs {
						close(sub)
						delete(subs, sub)
					}
				} else {
					for sub := range subs {
						select {
						case sub <- data:
						default:
						}
					}
				}
			}
		}
	}

	func (b *Broker) Subscribe() chan int {
		sub := make(chan int, 10)
		b.subChan <- sub
		return sub
	}

	func (b *Broker) Publish(data int) {
		b.dataChan <- data
	}

	func main() {
		b:=Init()
		var consumerWg sync.WaitGroup
		var producerWg sync.WaitGroup
		for i := range 5 {
			ch:=b.Subscribe()
			consumerWg.Add(1)
			go func() {
				defer consumerWg.Done()
				for msg:=range ch{
					fmt.Printf("consumer %d got number %d\n",i,msg)
				}
			}()
		}

		for i:=range 3{
			producerWg.Add(1)
			go func ()  {
				defer producerWg.Done()
				for j := range 3{
					time.Sleep(200*time.Millisecond)
					b.Publish(i*10+j)
				}
			}()
			time.Sleep(30*time.Millisecond)
		}
		go func ()  {
			producerWg.Wait()
			close(b.dataChan)
		}()
		consumerWg.Wait()
		fmt.Println("done")
	}