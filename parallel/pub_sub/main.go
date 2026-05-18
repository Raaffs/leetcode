package main

import (
	"fmt"
	"log"
	"sync"
)

type Broker struct {
	subChan   chan chan int
	unsubChan chan chan int
	pubChan   chan int
}

func NewBroker() *Broker {
	b := &Broker{
		subChan:   make(chan chan int),
		unsubChan: make(chan chan int),
		pubChan:   make(chan int),
	}
	go b.serve()
	return b
}

func (b *Broker) serve() {
	activeSubs := make(map[chan int]struct{})
	persist := make(map[chan int][]int)
	for {
		select {
		case ch := <-b.subChan:
			activeSubs[ch] = struct{}{}
		case data, ok := <-b.pubChan:
			if !ok {
				for sub := range activeSubs {
					if backlog, exists := persist[sub]; exists {
						for b := range backlog{
							sub<-b
						}
						close(sub)
					}else{
						close(sub)
					}
				}
				return 
			}
			for sub := range activeSubs {
				b.broadCast(data, sub, persist)
			}
		case sub := <-b.unsubChan:
			close(sub)
			delete(activeSubs, sub)
			delete(persist, sub)
		}
	}
}

func (b *Broker) Subscribe(channel chan int) { b.subChan <- channel }
func (b *Broker) Publish(data int)           { b.pubChan <- data }

func (b *Broker) broadCast(data int, sub chan int, persist map[chan int][]int) {
	if len(persist[sub]) > 0 {
		persist[sub] = append(persist[sub], data)
		b.tryDrainingBacklog(sub, persist)
		return
	}

	select {
	case sub <- data:
	default:
		persist[sub] = []int{data}
	}
}

func (b *Broker) tryDrainingBacklog(sub chan int, persist map[chan int][]int) {
	queue := persist[sub]
	for len(queue) > 0 {
		select {
		case sub <- queue[0]:
			queue = queue[1:]
		default:
			persist[sub] = queue
			return
		}
	}
	delete(persist, sub)
}

func main() {
	broker := NewBroker()
	var cWG sync.WaitGroup
	var pWG sync.WaitGroup
	for i := range 5 {
		cWG.Add(1)
		channel := make(chan int, 10)
		broker.Subscribe(channel)
		
		go func(id int) { 
			msgs:=[]int{}
			defer cWG.Done()
			for msg := range channel {
				fmt.Printf("subscribe %d got msg %d\n", id, msg)
				msgs=append(msgs, msg)
			}
			if len(msgs)!=30{
				log.Fatalf("expected 30 got %d\n",len(msgs))
			}else{
				fmt.Printf("subscribe %d received all messages successfully\n", id)
			}
		}(i + 1) 
	}

	for i := range 3 {
		pWG.Add(1)
		go func(pID int) {
			defer pWG.Done()
			for j := range 10 {
				broker.Publish(pID*10 + j)
			}
		}(i + 1)
	}

	go func() {
		pWG.Wait()
		close(broker.pubChan)
	}()

	cWG.Wait()
}