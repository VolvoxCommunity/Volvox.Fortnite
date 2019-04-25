package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/api"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"github.com/volvoxcommunity/volvox.fortnite/src/utils"
	"strings"
)

/**

 * Created by cxnky on 25/04/2019 at 00:56
 * commands
 * https://github.com/cxnky/

**/

func FetchWins(ctx framework.Context) {
	epicName := ""

	if len(ctx.Args) <= 1 {
		for _, m := range ctx.Guild.Members {
			if m.User.ID == ctx.User.ID {
				epicName = m.Nick
			}
		}

		if epicName == "" {
			epicName = ctx.User.Username
		}

		platform := ctx.GetUserPlatform()
		wins, err := api.FetchLifetimeWins(epicName, platform)

		if err != nil {
			ctx.ReplyErrorEmbed("```" + err.Error() + "```")
			return
		}

		ctx.ReplyEmbed(&discordgo.MessageEmbed{
			Description: "You have " + wins + " lifetime wins.",
			Color:       utils.GetInformationColour(),
		})

		return

	}

	platform := strings.ToLower(ctx.Args[0])

	if platform != "pc" && platform != "xbox" && platform != "psn" {
		ctx.Reply("Invalid option! Valid options are: pc/xbox/psn")
		return
	}

	epicName = strings.Title(strings.Join(strings.Fields(ctx.Message.Content)[2:], " "))
	wins, err := api.FetchLifetimeWins(epicName, platform)

	if err != nil {
		ctx.ReplyErrorEmbed("```" + err.Error() + "```")
		return
	}

	ctx.ReplyEmbed(&discordgo.MessageEmbed{
		Description: epicName + " has " + wins + " lifetime wins.",
		Color:       utils.GetInformationColour(),
	})

}
