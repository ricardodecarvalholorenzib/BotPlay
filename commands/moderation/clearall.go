package moderation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ClearAll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil || i.Member.Permissions&discordgo.PermissionAdministrator == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Você precisa de permissão de Administrador para usar este comando.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	confirm := i.ApplicationCommandData().Options[0].StringValue()

	if confirm != "confirm" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Você precisa digitar 'confirm' para excluir todas as mensagens.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	channelID := i.ChannelID
	messages, err := s.ChannelMessages(channelID, 100, "", "", "")
	if err != nil || len(messages) == 0 {
		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "❌ Nenhuma mensagem encontrada para excluir.",
		})
		return
	}

	var messageIDs []string
	for _, msg := range messages {
		messageIDs = append(messageIDs, msg.ID)
	}

	err = s.ChannelMessagesBulkDelete(channelID, messageIDs)
	if err != nil {
		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "❌ Erro ao apagar mensagens em massa (mensagens com mais de 14 dias não podem ser apagadas assim): " + err.Error(),
		})
		return
	}

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: fmt.Sprintf("✅ %d mensagens excluídas com sucesso!", len(messageIDs)),
	})
}