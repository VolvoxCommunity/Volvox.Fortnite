package main

import "C"
import (
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/api"
	"github.com/volvoxcommunity/volvox.fortnite/src/commands"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"github.com/volvoxcommunity/volvox.fortnite/src/logging"
	"os"
	"os/signal"
	"strconv"
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
	logging.InitLogger()

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
		panic("unable to establish a connection to Discord: " + err.Error())
	}

	logging.Log.Info("Volvox.Fortnite is now running. Ctrl+C to exit.")
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
		logging.Log.Error("Error getting channel,", err)
		return
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		logging.Log.Error("Error getting guild,", err)
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
	commandHandler.RegisterCommand("sync", commands.ForceTierSynchronisation)
	commandHandler.RegisterCommand("start", commands.StartSync)
}

func readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	logging.Log.Info("Volvox.Fortnite is ready! Gateway version " + strconv.Itoa(r.Version))
	err := s.UpdateStatus(0, "Volvox.Fortnite | Go v1.12")

	if err != nil {
		logging.Log.Info("unable to set status: " + err.Error())
	}

	api.APIKey = Config.TRNAPIKey

}
