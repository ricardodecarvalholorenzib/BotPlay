package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go-games/commands/general"
	"go-games/commands/games"
	"go-games/commands/moderation"

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

	var adminPermission int64 = discordgo.PermissionAdministrator

	commands := []*discordgo.ApplicationCommand{

		{
			Name:        "msg",
			DefaultMemberPermissions: &adminPermission,
			Description: "Envia uma mensagem com o Bot",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "mensagem",
					Description: "Mensagem a ser enviada",
					Required:    true,
				},
			},
		},

		{ 
			Name: 	      "clear",
			Description:  "Exclui a mensagem (usando seu ID)",
			DefaultMemberPermissions:  &adminPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: 		discordgo.ApplicationCommandOptionString,
					Name: 		"id",
					Description: 	"ID da mensagem a ser excluída",
					Required: 	true,
				},
			},
		},

		{
			Name:        "clearall",
			Description: "Exclui todas as mensagens do canal",
			DefaultMemberPermissions: &adminPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "confirm",
					Description: "Confirmação para excluir todas as mensagens",
					Required:    true,
				},
			},
		},

		{
			Name:        "kick",
			Description: "Expulsa um usuário do servidor",
			DefaultMemberPermissions: &adminPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "usuário",
					Description: "ID do Usuário a ser expulso",
					Required:    true,
				},
			},
		},

		{
			Name:        "caraoucoroa",
			Description: "Jogue cara ou coroa contra o bot!",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "escolha",
					Description: "Escolha entre cara ou coroa",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Cara 🪙",
							Value: "cara",
						},
						{
							Name:  "Coroa 👑",
							Value: "coroa",
						},
					},
				},
			},
		},

		{
    		Name:        "unban",
    		Description: "Desbane um usuário do servidor",
			DefaultMemberPermissions: &adminPermission,
    		Options: []*discordgo.ApplicationCommandOption{
        		{
            		Type:        discordgo.ApplicationCommandOptionString,
            		Name:        "id",
					Description: "ID do Usuário a ser desbanido",
					Required:    true,
				},
			},
		},

	}

	guildID := "1544872649573007400"

	for _, command := range commands {
		_, err := bot.ApplicationCommandCreate(
			bot.State.User.ID,
			guildID,
			command,
		)

		if err != nil {
			fmt.Println("Erro ao registrar /" + command.Name + ":", err)
		}
	}

	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		data := i.ApplicationCommandData()

		switch data.Name {

		case "ban":
			moderation.Ban(s, i)

		case "ping":
			general.Ping(s, i)

		case "dado":
			games.Dado(s, i)

		case "help":
			general.Help(s, i)

		case "unban":
			moderation.Unban(s, i)

		case "clear":
			moderation.Clear(s, i)

		case "kick":
			moderation.Kick(s, i)

		case "clearall":
			moderation.ClearAll(s, i)

		case "caraoucoroa":
			games.CaraOuCoroa(s, i)

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