package framework

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/utils"
	"strings"
)

/**

 * Created by cxnky on 24/04/2019 at 15:45
 * framework
 * https://github.com/cxnky/

**/

// Context is the struct which contains all of the necessary information for a command
type Context struct {
	Discord      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Message      *discordgo.MessageCreate
	Args         []string
	Config       Configuration

	CmdHandler *CommandHandler
}

// NewContext creates an instance of the Context struct and populates it with data to be passed to the command
func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel,
	user *discordgo.User, message *discordgo.MessageCreate, cmdHandler *CommandHandler, config Configuration) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.CmdHandler = cmdHandler
	ctx.Config = config

	return ctx
}

// Reply responds to the user in the form @user, <message>
func (ctx Context) Reply(content ...string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, strings.Join(content, " "))
	if err != nil {
		fmt.Println("Unable to send message:", err)
		return nil
	}
	return msg
}

// UserHasRole checks whether or not the user has a specific role
func (ctx Context) UserHasRole(roleID string) bool {
	users := ctx.Guild.Members

	for _, u := range users {
		if u.User.ID == ctx.User.ID {
			for _, r := range u.Roles {
				if r == roleID {
					return true
				}
			}
		}
	}
	return false
}

// ReplyErrorEmbed replies to the user in the form of an error embed (red and with title "Error")
func (ctx Context) ReplyErrorEmbed(description string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, &discordgo.MessageEmbed{
		Title:       "Error",
		Description: description,
		Color:       utils.GetErrorColour(),
	})
	if err != nil {
		fmt.Println("Error whilst sending embed:", err)
		return nil
	}
	return msg
}

// GetUserPlatform checks whether the user has any of the platform roles
func (ctx Context) GetUserPlatform() string {
	if ctx.UserHasRole(ctx.Config.XboxRole) {
		return "xbox"
	}
	if ctx.UserHasRole(ctx.Config.PCRole) {
		return "pc"
	}
	if ctx.UserHasRole(ctx.Config.PS4Role) {
		return "ps4"
	}

	return "ns"

}

// ReplyEmbed allows you to reply to the user using a custom embed that is defined
func (ctx Context) ReplyEmbed(embed *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed)
	if err != nil {
		fmt.Println("Unable to send embed:", err)
		return nil
	}
	return msg
}
