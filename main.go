package main

import (
	"fmt"
	"github.com/efimshifrin/discobot/apiclients"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	apiclients.DiscordClient()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
