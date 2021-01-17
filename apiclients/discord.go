package apiclients

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/efimshifrin/discobot/handlers"
	"github.com/efimshifrin/discobot/utils"
)

func DiscordClient() {

	dg, err := discordgo.New("Bot " + utils.GetEnvString("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(handlers.MessageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

}
