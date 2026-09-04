package moderation

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func Kick(s *discordgo.Session, i *discordgo.InteractionCreate) {

	user := i.ApplicationCommandData().Options[0].UserValue(s)

	err := s.GuildMemberDelete(i.GuildID, user.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Erro ao expulsar o usuário: " + err.Error(),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("✅ Usuário `%s` expulso com sucesso!", user.Username),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		},
	)
}