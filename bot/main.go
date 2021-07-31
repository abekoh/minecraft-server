package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	stopBot = make(chan bool)
)

func main() {
	discordOperator, err := newMinecraftServerDiscordOperator(os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_BOT_CLIENT_ID"), GoogleComputeEngineOperator{})
	if err != nil {
		panic(err)
	}

	err = discordOperator.open()
	if err != nil {
		panic(err)
	}
	defer discordOperator.close()
	fmt.Println("Listening...")
	<-stopBot
}

type DiscordOperator interface {
	open() error
	close() error
}

type MinecraftServerDiscordOperator struct {
	discordSession *discordgo.Session
}

func newMinecraftServerDiscordOperator(botToken string, botClientId string, serverOperator ServerOperator) (MinecraftServerDiscordOperator, error) {
	if len(botToken) == 0 {
		return MinecraftServerDiscordOperator{}, fmt.Errorf("token is not defined")
	}
	if len(botClientId) == 0 {
		return MinecraftServerDiscordOperator{}, fmt.Errorf("bot client-id is not defined")
	}
	discord, err := discordgo.New(botToken)
	if err != nil {
		return MinecraftServerDiscordOperator{}, err
	}
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		botMentionPattern := regexp.MustCompile(fmt.Sprintf("<@%s> (.*)", botClientId))
		message := botMentionPattern.ReplaceAllString(m.Content, "$1")
		if len(message) == 0 {
			return
		}
		if strings.Contains(message, "start") {
			serverOperator.start()
		}
		if strings.Contains(message, "stop") {
			serverOperator.stop()
		}
	})
	return MinecraftServerDiscordOperator{
		discordSession: discord,
	}, nil
}

func (msdo MinecraftServerDiscordOperator) open() error {
	return msdo.discordSession.Open()
}

func (msdo MinecraftServerDiscordOperator) close() error {
	return msdo.discordSession.Close()
}

type ServerOperator interface {
	start() error
	stop() error
}

type GoogleComputeEngineOperator struct {
}

func (gceo GoogleComputeEngineOperator) start() error {
	fmt.Print("start server")
	return nil
}

func (gceo GoogleComputeEngineOperator) stop() error {
	fmt.Print("stop server")
	return nil
}
