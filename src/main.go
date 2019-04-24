package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

/**

 * Created by cxnky on 24/04/2019 at 15:20
 * main
 * https://github.com/cxnky/

**/

var Config Configuration

func main() {
	Config = ReadConfig()
	dg, err := discordgo.New("Bot " + Config.Token)

	if err != nil {
		panic(err)
	}

	dg.AddHandler(readyHandler)
	err = dg.Open()

	if err != nil {
		panic("unable to establish connection to Discord: " + err.Error())
	}

	fmt.Println("Volvox.Fortnite is now running. Ctrl+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()

}

func readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Volvox.Fortnite is ready! Gateway version", r.Version)
	err := s.UpdateStatus(0, "Volvox.Fortnite | Go v1.11")

	if err != nil {
		_, _ = fmt.Println(fmt.Errorf("unable to set bot status: " + err.Error()))
	}

}
