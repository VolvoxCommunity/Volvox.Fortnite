package commands

import (
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"github.com/volvoxcommunity/volvox.fortnite/src/services"
)

func StartSync(ctx framework.Context) {
	if ctx.UserHasRole("568827043559768069") || ctx.UserHasRole("570448844224200735") {
		go services.RunSyncService(ctx)
		ctx.Reply("Sync service started.")
	} else {
		ctx.ReplyErrorEmbed("You must be a moderator or higher to use this command!")
	}
}
