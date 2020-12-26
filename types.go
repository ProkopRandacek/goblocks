package main

import (
	"os/exec"
	"bytes"
	"time"
	"strings"
	"os"
)

type Block struct {
	Command string
	Interval time.Duration
	Separator string
	Color string
	Value string
}

type content struct {
	blocks *[]Block
	value string
}

func (c *content) Update() {
	c.value = string((*c.blocks)[len((*c.blocks))-1].Separator[len((*c.blocks)[len((*c.blocks))-1].Separator)-1])
	for i := range *c.blocks {
		c.value += " "
		c.value += (*c.blocks)[i].Value
		c.value += " "
		if i != len(*c.blocks)-1 {
			c.value += (*c.blocks)[i].Separator
		}
	}
	c.value += string((*c.blocks)[0].Separator[0])
	exec.Command("xsetroot", "-name", c.value).Run()
} 

func (b *Block) Update() {
	cmd := exec.Command("./" + b.Command)
	cmd.Dir = os.Getenv("HOME") + "/.config/goblocks/"
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	b.Value = strings.Split(out.String(), "\n")[0]
}
