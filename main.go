package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	sess, err := discordgo.New("Bot MTIxMzc2NTY4MjE2OTM4NDk4MA.GWo6IL.RDifTo4-eTdY30mLaLNbgaCoV_s8VH_IVene3Y" /*your token*/)
	if err != nil {
		log.Fatal(err)

	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "Hello" {
			s.ChannelMessageSend(m.ChannelID, "World !")
		}

	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The Bot is Online ! ")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
