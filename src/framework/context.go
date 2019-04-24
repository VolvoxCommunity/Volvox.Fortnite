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

type Context struct {
	Discord      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Message      *discordgo.MessageCreate
	Args         []string

	CmdHandler *CommandHandler
}

func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel,
	user *discordgo.User, message *discordgo.MessageCreate, cmdHandler *CommandHandler) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.CmdHandler = cmdHandler

	return ctx
}

func (ctx Context) Reply(content ...string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, strings.Join(content, " "))
	if err != nil {
		fmt.Println("Unable to send message:", err)
		return nil
	}
	return msg
}

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

func (ctx Context) ReplyEmbed(embed *discordgo.MessageEmbed) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed)
	if err != nil {
		fmt.Println("Unable to send embed:", err)
		return nil
	}
	return msg
}
