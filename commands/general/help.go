package general

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "📋 Aqui estão os comandos disponíveis:\n" +
					 "\n- /ping: Responde com Pong!" + 
					 "\n- /help: Exibe esta mensagem de ajuda" +
					 "\n- /dado: Rola um dado de 1 a 6",
		},
	})
	if err != nil {
		return
	}
}