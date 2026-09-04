package games

import (
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func Dado(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resultado := rand.IntN(6) + 1

	var mensagem string

	switch resultado {
	case 1:
		mensagem = "🎲 Você rolou um dado e caiu no número 1! Que azar! 😢"
	case 2:
		mensagem = "🎲 Você rolou um dado e caiu no número 2! Foi quase. Muito melhor que 1!"
	case 3:
		mensagem = "🎲 Você rolou um dado e caiu no número 3! Não está ruim!"
	case 4:
		mensagem = "🎲 Você rolou um dado e caiu no número 4! Está indo bem!"
	case 5:
		mensagem = "🎲 Você rolou um dado e caiu no número 5! Muito bom! Mas não é um 6! 😂"
	case 6:
		mensagem = "🎲 Você rolou um dado e caiu no número 6! Que sorte! 😎"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: mensagem,
		},
	})
}