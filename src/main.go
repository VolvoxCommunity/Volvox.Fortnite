package main

import "C"
import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/api"
	"github.com/volvoxcommunity/volvox.fortnite/src/commands"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

/**

 * Created by cxnky on 24/04/2019 at 15:20
 * main
 * https://github.com/cxnky/

**/

var (
	// Config is the object in which the parsed configuration will be stored
	Config         framework.Configuration
	commandHandler *framework.CommandHandler
)

func main() {
	Config = framework.ReadConfig()
	dg, err := discordgo.New("Bot " + Config.Token)
	commandHandler = framework.NewCommandHandler()

	if err != nil {
		panic(err)
	}

	registerCommands()

	dg.AddHandler(readyHandler)
	dg.AddHandler(messageReceived)

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

func messageReceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author
	if user.ID == Config.ClientID || user.Bot {
		return
	}

	content := m.Content
	if len(content) < 1 {
		return
	}

	if content[:len(Config.Prefix)] != Config.Prefix {
		return
	}
	content = content[len(Config.Prefix):]
	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := commandHandler.Get(name)
	if !found {
		return
	}
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return
	}
	ctx := framework.NewContext(s, guild, channel, user, m, commandHandler, Config)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)

}

func registerCommands() {
	commandHandler.RegisterCommand("ping", commands.PingCommand)
	commandHandler.RegisterCommand("wins", commands.FetchWins)
}

func readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Volvox.Fortnite is ready! Gateway version", r.Version)
	err := s.UpdateStatus(0, "Volvox.Fortnite | Go v1.11")

	if err != nil {
		_, _ = fmt.Println(fmt.Errorf("unable to set bot status: " + err.Error()))
	}

	api.APIKey = Config.TRNAPIKey

}
