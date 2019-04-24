package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/volvoxcommunity/volvox.fortnite/src/framework"
	"github.com/volvoxcommunity/volvox.fortnite/src/utils"
	"time"
)

/**

 * Created by cxnky on 24/04/2019 at 15:52
 * commands
 * https://github.com/cxnky/

**/

func PingCommand(ctx framework.Context) {
	latency := ctx.Discord.HeartbeatLatency().Round(time.Millisecond).String()

	ctx.ReplyEmbed(&discordgo.MessageEmbed{
		Description: ":ping_pong: " + latency,
		Color:       utils.GetInformationColour(),
	})
}
