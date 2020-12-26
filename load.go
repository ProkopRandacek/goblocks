package main

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

func Load() []Block {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.config/goblocks/blocks.json")
	if err != nil {
		panic(err)
	}
	var toReturn []Block
	err = json.Unmarshal(dat, &toReturn)
	if err != nil {
		panic(err)
	}
	return toReturn
}
