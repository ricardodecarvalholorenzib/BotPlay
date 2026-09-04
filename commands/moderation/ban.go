package moderation

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func Ban(s *discordgo.Session, i *discordgo.InteractionCreate) {
	username := i.ApplicationCommandData().Options[0].UserValue(s).Username

	user := i.ApplicationCommandData().Options[0].UserValue(s)

	err := s.GuildBanCreateWithReason(i.GuildID, user.ID, "Banned by bot command", 0)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Erro ao banir o usuário: " + err.Error(),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("✅ Usuário com ID `%s` banido com sucesso!", username),
		},
	})
}