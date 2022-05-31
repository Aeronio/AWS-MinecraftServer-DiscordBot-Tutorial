package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters

var (
	token = "yourDiscordBotToken" //Your Discord Bot private token
	//cognitoToken = `Your Cognito credentials token` // Your cognito token
)

func ConnectToDiscord() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from bot himself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!mcstart" {
		s.ChannelMessageSend(m.ChannelID, "Checking Server...")
		sb := startServer()
		s.ChannelMessageSend(m.ChannelID, sb)
		return
	}

	if m.Content == "!mcstop" {
		s.ChannelMessageSend(m.ChannelID, "Checking Server...")
		sb := stopServer()
		s.ChannelMessageSend(m.ChannelID, sb)
		return
	}

	if m.Content == "!mcstatus" {
		s.ChannelMessageSend(m.ChannelID, "Checking Server...")
		sb := statusServer()
		s.ChannelMessageSend(m.ChannelID, sb)
		return
	}

	if m.Content == "!mcreboot" {
		s.ChannelMessageSend(m.ChannelID, "Checking Server...")
		sb := rebootServer()
		s.ChannelMessageSend(m.ChannelID, sb)
		return
	}

	if m.Content == "!mchelp" {
		handleHelp(s, m)
		return
	}

	handleHelp(s, m)
}

func startServer() string {

	url := "https://yourAPIPrefix.execute-api.your-amazon-region.amazonaws.com/dev/start"

	client := &http.Client{}

	req := myAuth(url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
	return sb
}

func stopServer() string {

	url := "https://yourAPIPrefix.execute-api.your-amazon-region.amazonaws.com/dev/stop"

	client := &http.Client{}

	req := myAuth(url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
	return sb
}

func statusServer() string {

	url := "https://yourAPIPrefix.execute-api.your-amazon-region.amazonaws.com/dev/status"

	client := &http.Client{}

	req := myAuth(url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
	return sb
}

func rebootServer() string {

	url := "https://yourAPIPrefix.execute-api.your-amazon-region.amazonaws.com/dev/reboot"

	client := &http.Client{}

	req := myAuth(url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
	return sb
}

func myAuth(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("authorizationToken", `abc123`)
	//req.Header.Add("Authorization", cognitoToken)
	return req
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := "commands:\n"
	msg = fmt.Sprintf("%s\t%s\n", msg, "!mchelp - shows all available commands")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!mcstart - starts your Minecraft server")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!mcstop - stops your Minecraft server")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!mcreboot - reboots your Minecraft server")
	msg = fmt.Sprintf("%s\t%s\n", msg, "!mcstatus - shows when the Minecraft server was started")

	s.ChannelMessageSend(m.ChannelID, msg)
}
