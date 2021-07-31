package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

var (
	stopBot = make(chan bool)
)

func main() {
	var project, zone, instanceName, credPath string
	flag.StringVar(&project, "project", "", "GCP project")
	flag.StringVar(&zone, "zone", "", "GCP zone")
	flag.StringVar(&instanceName, "name", "", "GCP instance name")
	flag.StringVar(&credPath, "serviceaccount", "", "GCP service account secret path")
	computeEngineOperator, err := newGoogleComputeEngineOperator(credPath, project, zone, instanceName)
	if err != nil {
		panic(err)
	}

	discordOperator, err := newMinecraftServerDiscordOperator(
		os.Getenv("DISCORD_TOKEN"),
		os.Getenv("DISCORD_BOT_CLIENT_ID"),
		computeEngineOperator,
	)
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
		if strings.Contains(message, "wakeup") {
			serverOperator.start()
		}
		if strings.Contains(message, "shutdown") {
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
	start()
	stop()
}

type GoogleComputeEngineOperator struct {
	service      *compute.Service
	project      string
	zone         string
	instanceName string
}

func newGoogleComputeEngineOperator(creadentialFilePath string, project string, zone string, instanceName string) (GoogleComputeEngineOperator, error) {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		return GoogleComputeEngineOperator{}, err
	}
	service, err := compute.New(client)
	if err != nil {
		return GoogleComputeEngineOperator{}, err
	}
	return GoogleComputeEngineOperator{
		service:      service,
		project:      project,
		zone:         zone,
		instanceName: instanceName,
	}, nil
}

func (gceo GoogleComputeEngineOperator) start() {
	fmt.Print("start server")
}

func (gceo GoogleComputeEngineOperator) stop() {
	fmt.Print("stop server")
}
