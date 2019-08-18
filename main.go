package main

import (
	"calendarproto/cmd"
	"log"
)

var (
	version = "dev"
	build   = "local"
)

func main() {
	log.Printf("Started calendarproto %s-%s", version, build)
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
