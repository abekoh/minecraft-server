package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	compute "google.golang.org/api/compute/v1"
)

var (
	stopBot = make(chan bool)
)

func main() {
	var project, zone, instanceName string
	flag.StringVar(&project, "project", "", "GCP project")
	flag.StringVar(&zone, "zone", "", "GCP zone")
	flag.StringVar(&instanceName, "name", "", "GCP instance name")
	flag.Parse()
	computeEngineOperator, err := newGoogleComputeEngineOperator(project, zone, instanceName)
	if err != nil {
		panic(err)
	}

	discordOperator, err := newMinecraftServerDiscordOperator(
		os.Getenv("DISCORD_TOKEN"),
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

func newMinecraftServerDiscordOperator(botToken string, serverOperator ServerOperator) (MinecraftServerDiscordOperator, error) {
	if len(botToken) == 0 {
		return MinecraftServerDiscordOperator{}, fmt.Errorf("token is not defined")
	}
	discord, err := discordgo.New(botToken)
	if err != nil {
		return MinecraftServerDiscordOperator{}, err
	}
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		botMentionPattern := regexp.MustCompile(fmt.Sprintf("<@%s> (.*)", s.State.User.ID))
		message := botMentionPattern.ReplaceAllString(m.Content, "$1")
		if len(message) == 0 {
			return
		}
		if strings.Contains(message, "wakeup") {
			err := serverOperator.wakeup()
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> failed to wakeup: %v", m.Author.ID, err))
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> succeeded to wakeup server!", m.Author.ID))
			}
			return
		}
		if strings.Contains(message, "shutdown") {
			err := serverOperator.shutdown()
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> failed to shutdown: %v", m.Author.ID, err))
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> succeeded to shutdown server!", m.Author.ID))
			}
			return
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
	wakeup() error
	shutdown() error
}

type GoogleComputeEngineOperator struct {
	service      *compute.Service
	project      string
	zone         string
	instanceName string
}

func newGoogleComputeEngineOperator(project string, zone string, instanceName string) (GoogleComputeEngineOperator, error) {
	ctx := context.Background()
	// client, err := google.DefaultClient(ctx, compute.ComputeScope)
	// if err != nil {
	// 	return GoogleComputeEngineOperator{}, err
	// }
	service, err := compute.NewService(ctx)
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

func (gceo GoogleComputeEngineOperator) wakeup() error {
	fmt.Println("wake-upping server...")
	instance, err := gceo.service.Instances.Get(gceo.project, gceo.zone, gceo.instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to get instance: %v", err)
	}
	if instance.Status != "TERMINATED" && instance.Status != "STOPPED" && instance.Status != "SUSPENDED" {
		return fmt.Errorf("instance's status is not TERMINATED/STOPPED/SUSPENDED")
	}
	_, err = gceo.service.Instances.Start(gceo.project, gceo.zone, gceo.instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to start instance: %v", err)
	}
	fmt.Println("succeded to wake-up server!")
	return nil
}

func (gceo GoogleComputeEngineOperator) shutdown() error {
	fmt.Println("shutdowning server...")
	instance, err := gceo.service.Instances.Get(gceo.project, gceo.zone, gceo.instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to get instance: %v", err)
	}
	if instance.Status != "RUNNING" {
		return fmt.Errorf("instance's status is not RUNNING")
	}
	_, err = gceo.service.Instances.Stop(gceo.project, gceo.zone, gceo.instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to stop instance: %v", err)
	}
	fmt.Println("succeded to shutdown server!")
	return nil
}
