# 🎮 BotPlay

> 🤖 Bot de Discord desenvolvido em **Go** para reunir jogos, comandos e diversão em um único lugar.

![Go](https://img.shields.io/badge/Go-1.27+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Discord](https://img.shields.io/badge/Discord-Bot-5865F2?style=for-the-badge&logo=discord&logoColor=white)
![DiscordGo](https://img.shields.io/badge/DiscordGo-API%20Wrapper-7289DA?style=for-the-badge)

## 🧩 Sobre o projeto

O **BotPlay** é um projeto pessoal criado para aprender **Go** na prática através do desenvolvimento de um bot para Discord.

A ideia é transformar o projeto em um espaço com pequenos jogos e sistemas interativos, enquanto exploro recursos da linguagem, APIs, eventos e armazenamento de dados.

## 🚀 Atualmente

O bot já possui:

- `/ping` — verifica se o bot está online
- `/help` — exibe os comandos disponíveis
- ⚡ Slash Commands do Discord
- 🔐 Token configurado por variável de ambiente
- 🔌 Conexão com o Discord através da biblioteca DiscordGo

## 🎮 Próximos passos

- [ ] 🎲 Jogo de dados
- [ ] ♋ Cara ou coroa
- [ ] 🧠 Adivinhe o número
- [ ] ✊ Pedra, Papel e Tesoura
- [ ] 🏆 Sistema de XP e ranking
- [ ] 💰 Sistema de moedas
- [ ] 💾 Banco de dados
- [ ] 🎖️ Conquistas

## 🛠️ Tecnologias

- **Go** — linguagem principal
- **DiscordGo** — integração com a API do Discord
- **Discord Slash Commands** — sistema de comandos

## 📁 Estrutura

```text
BotPlay/
└── Bot-Go/
    ├── main.go
    ├── go.mod
    └── go.sum
```

## ▶️ Executando localmente

### 1. Clone o projeto

```bash
git clone https://github.com/ricardodecarvalholorenzib/BotPlay.git
cd BotPlay/Bot-Go
```

### 2. Configure o token

O token do bot não deve ser colocado diretamente no código.

No PowerShell:

```powershell
$env:DISCORD_BOT_TOKEN = "SEU_TOKEN"
```

### 3. Execute

```bash
go run .
```

Se tudo estiver correto, o bot será conectado ao Discord e os Slash Commands serão registrados.

## 🔒 Segurança

**Nunca publique o token do bot no GitHub.**

O projeto utiliza a variável de ambiente `DISCORD_BOT_TOKEN` para manter a credencial fora do código-fonte. Arquivos `.env` também estão incluídos no `.gitignore`.

Se um token for exposto, ele deve ser redefinido no Discord Developer Portal imediatamente.

## 📌 Objetivo

Mais do que criar um bot, este projeto serve como laboratório para aprender **Go construindo algo real**.

Cada nova funcionalidade será uma oportunidade para aprender uma parte diferente da linguagem e do ecossistema Go.

---

⭐ Projeto em desenvolvimento por [Ricardo de Carvalho](https://github.com/ricardodecarvalholorenzib)
