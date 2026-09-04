package moderation

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Unban(s *discordgo.Session, i *discordgo.InteractionCreate) {
	targetID := i.ApplicationCommandData().Options[0].StringValue()

	err := s.GuildBanDelete(i.GuildID, targetID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Erro ao desbanir o usuário: " + err.Error(),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("✅ Usuário com ID `%s` desbanido com sucesso!", targetID),
		},
	})
}