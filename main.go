package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")

	fmt.Println("Token encontrado?", token != "")

	if token == "" {
		fmt.Println("Please set the DISCORD_BOT_TOKEN environment variable.")
		return
	}

	bot, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	err = bot.Open()

	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("🤖 Bot conectado!")

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Responde com Pong!",
		},
		{
			Name:        "help",
			Description: "Mostra os comandos disponíveis",
		},
	}

	for _, command := range commands {
		_, err := bot.ApplicationCommandCreate(
			bot.State.User.ID,
			"",
			command,
		)

		if err != nil {
			fmt.Println("Erro ao registrar /" + command.Name + ":", err)
		}
	}

	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		data := i.ApplicationCommandData()

		switch data.Name {

		case "ping":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "🏓 Pong!",
				},
			})

		case "help":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "🎮 **Comandos disponíveis:**\n\n/ping - Responde com Pong!\n/help - Mostra esta mensagem",
				},
			})
		}
	})

	fmt.Println("🎮 Slash Commands registrados!")
	fmt.Println("Pressione CTRL+C para sair.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	<-sc

	fmt.Println("Shutting down...")

	bot.Close()
}