package moderation

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func Clear(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := i.ApplicationCommandData().Options[0].StringValue()
	channelID := i.ChannelID

	err := s.ChannelMessageDelete(channelID, id)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Erro ao excluir a mensagem: " + err.Error(),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			 	Content: fmt.Sprintf("✅ Mensagem com ID `%s` excluída com sucesso!", id),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		},

	)

}