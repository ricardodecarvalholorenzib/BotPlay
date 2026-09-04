package games

import (
	"fmt"
	"math/rand/v2" // Se estiver usando Go 1.22+

	"github.com/bwmarrin/discordgo"
)

func CaraOuCoroa(s *discordgo.Session, i *discordgo.InteractionCreate) {
	escolhaUsuario := i.ApplicationCommandData().Options[0].StringValue()

	if escolhaUsuario != "cara" && escolhaUsuario != "coroa" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Opção inválida! Escolha 'cara' ou 'coroa'.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	sorteio := rand.IntN(2) 

	var resultadoSorteio string
	if sorteio == 0 {
		resultadoSorteio = "cara"
	} else {
		resultadoSorteio = "coroa"
	}

	var mensagem string
	if escolhaUsuario == resultadoSorteio {
		mensagem = fmt.Sprintf("Você escolheu: ``%s``!\n Deu **%s**! Você **ganhou**! 🎉", escolhaUsuario, resultadoSorteio)
	} else {
		mensagem = fmt.Sprintf("Você escolheu: ``%s``!\n Deu **%s**! Você **perdeu**! 😢", escolhaUsuario, resultadoSorteio)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: mensagem,
		},
	})
}