package main

import (
	"time"

	"github.com/lisieckir/ping-machine/network"
	"github.com/lisieckir/ping-machine/ui"
)

var data []int64
var lastChanged = make(chan int64)

func main() {
	a := ui.Ui{}
	go func() { a.Init() }()

	for {
		go async()

		select {
		case <-lastChanged:
			a.Draw(data)
		}
		time.Sleep(1 * time.Second)
	}
}

func async() {
	host1Response, _ := network.Handle("wp.pl")
	data = append(data, host1Response)
	now := time.Now().Unix()
	lastChanged <- now
}
