package services

import (
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/api"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"github.com/volvoxcommunity/volvox.fortnite/src/logging"
	"strconv"
	"time"
)

// RunSyncService starts the synchronisation process of the user tiers, which will be run every hour from initial boot
func RunSyncService(ctx framework.Context) {
	logChannel, err := ctx.Discord.Channel(ctx.Config.LogChannel)

	if err != nil {
		logging.Log.Error("could not fetch log channel: " + err.Error() + " aborting sync service start")
		return
	}

	for {
		for _, u := range ctx.Guild.Members {
			if !u.User.Bot {
				tier := fetchUpdatedTier(u, ctx)

				if tier == -1 {
					logging.Log.Error("error whilst fetching the stats for " + u.User.Username)
				} else {
					currentTier := fetchCurrentTier(u, ctx)
					if tier == currentTier {
						logging.Log.Info(u.User.Username + " is still at the same tier as they were previously")
					} else if currentTier == -1 {
						logging.Log.Info(u.User.Username + " should be given a Tier role")

						err := ctx.Discord.GuildMemberRoleAdd(ctx.Guild.ID, u.User.ID, fetchRoleByTier(tier, ctx))

						if err != nil {
							logging.Log.Error(err.Error())
							continue
						}

					} else if tier < currentTier {
						logging.Log.Info(u.User.Username + " should go up a tier!")
						currentRole := fetchRoleByTier(currentTier, ctx)
						newRole := fetchRoleByTier(tier, ctx)

						err := ctx.Discord.GuildMemberRoleRemove(ctx.Guild.ID, u.User.ID, currentRole)

						if err != nil {
							logging.Log.Error(err.Error())
							continue
						}

						err = ctx.Discord.GuildMemberRoleAdd(ctx.Guild.ID, u.User.ID, newRole)

						if err != nil {
							logging.Log.Error(err.Error())
							continue
						}

						_, _ = ctx.Discord.ChannelMessageSendEmbed(logChannel.ID, &discordgo.MessageEmbed{
							Title:       "Tier up!",
							Description: u.User.Mention() + " has tiered up from Tier " + strconv.Itoa(currentTier) + " to Tier " + strconv.Itoa(tier) + "!",
							Color:       0x00ff00,
						})

					}
				}

			}
		}

		time.Sleep(1 * time.Hour)
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
