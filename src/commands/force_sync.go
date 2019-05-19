package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/api"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"strconv"
)

/**

 * Created by cxnky on 25/04/2019 at 12:13
 * commands
 * https://github.com/cxnky/

**/

// ForceTierSynchronisation forces a synchronisation of all of the tiers of those who are registered
func ForceTierSynchronisation(ctx framework.Context) {

	if ctx.UserHasRole("568827043559768069") || ctx.UserHasRole("570448844224200735") {
		message, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, "Synchronising user tiers with their wins...")

		if err != nil {
			fmt.Println(fmt.Errorf("unable to send message: " + err.Error()))
			return
		}

		promoted := 0
		errors := 0

		for _, u := range ctx.Guild.Members {
			if !u.User.Bot {
				tier := fetchUpdatedTier(u, ctx)

				if tier == -1 {
					errors++
				} else {
					currentTier := fetchCurrentTier(u, ctx)
					if tier == currentTier {
						fmt.Println(u.User.Username, "is still at the same tier as they were previously")
					} else if currentTier == -1 {
						fmt.Println(u.User.Username, "should be given a Tier role")

						err := ctx.Discord.GuildMemberRoleAdd(ctx.Guild.ID, u.User.ID, fetchRoleByTier(tier, ctx))

						if err != nil {
							errors++
							fmt.Println(err.Error())
						}

					} else if tier < currentTier {
						fmt.Println(u.User.Username, "should go up a tier!")
						promoted++
					}
				}

			}
		}

		_, _ = ctx.Discord.ChannelMessageEdit(ctx.TextChannel.ID, message.ID, "Finished synchronising.\nPromoted: "+strconv.Itoa(promoted)+"\nUnregistered: "+strconv.Itoa(errors))
	} else {
		ctx.ReplyErrorEmbed("Only moderators and higher can use this command!")
	}

}

func fetchRoleByTier(tier int, ctx framework.Context) string {
	if tier == 5 {
		return ctx.Config.Roles.Tier5
	} else if tier == 4 {
		return ctx.Config.Roles.Tier4
	} else if tier == 3 {
		return ctx.Config.Roles.Tier3
	} else if tier == 2 {
		return ctx.Config.Roles.Tier2
	} else if tier == 1 {
		return ctx.Config.Roles.Tier1
	}

	return "-1"
}

func fetchCurrentTier(m *discordgo.Member, ctx framework.Context) int {
	ctx.User = m.User

	if ctx.UserHasRole(ctx.Config.Roles.Tier1) {
		return 1
	} else if ctx.UserHasRole(ctx.Config.Roles.Tier2) {
		return 2
	} else if ctx.UserHasRole(ctx.Config.Roles.Tier3) {
		return 3
	} else if ctx.UserHasRole(ctx.Config.Roles.Tier4) {
		return 4
	} else if ctx.UserHasRole(ctx.Config.Roles.Tier5) {
		return 5
	} else {
		return -1
	}
}

// This will return the role ID of the tier that the user should be on
func fetchUpdatedTier(m *discordgo.Member, ctx framework.Context) int {
	epicName := ""
	ctx.User = m.User

	if m.Nick == "" {
		epicName = m.User.Username
	} else {
		epicName = m.Nick
	}

	wins, err := api.FetchLifetimeWins(epicName, ctx.GetUserPlatform())

	if err != nil {
		return -1
	}

	winsInt, err := strconv.Atoi(wins)

	if err != nil {
		return -1
	}

	if winsInt >= 0 && winsInt <= 100 {
		return 5
	} else if winsInt >= 101 && winsInt <= 250 {
		return 4
	} else if winsInt >= 251 && winsInt <= 500 {
		return 3
	} else if winsInt >= 501 && winsInt <= 750 {
		return 2
	} else if winsInt >= 751 {
		return 1
	}

	return -1

}
