package commands

import (
	"github.com/volvoxcommunity/volvox.fortnite/src/api"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"strings"
)

/**

 * Created by cxnky on 25/04/2019 at 00:56
 * commands
 * https://github.com/cxnky/

**/

func FetchWins(ctx framework.Context) {
	if len(ctx.Args) <= 1 {
		ctx.Reply("Invalid arguments! Example: ``?stats pc connorw``")
		return
	}

	platform := strings.ToLower(ctx.Args[0])

	if platform != "pc" && platform != "xbox" && platform != "psn" {
		ctx.Reply("Invalid option! Valid options are: pc/xbox/psn")
		return
	}

	epicName := strings.Join(strings.Fields(ctx.Message.Content)[2:], " ")
	wins, err := api.FetchLifetimeWins(epicName, platform)

	if err != nil {
		ctx.Reply("An error occurred.\n```\n" + err.Error() + "\n```")
		return
	}

	ctx.Reply("You have " + wins + " lifetime wins.")

}
