package main

import (
	"time"
)

func run(b Block, c *content, pos int) {
	for {
		b.Update()
		(*c.blocks)[pos] = b
		c.Update()
		time.Sleep(b.Interval * time.Second)
	}
}
