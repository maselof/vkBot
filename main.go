package main

import (
	"flag"
	"log"
	"vkBot/clients/botsLP"
	"vkBot/clients/vkAPI"
	"vkBot/consumer/eventConsumer"
	vkBot2 "vkBot/events/vkBot"
)

func main() {
	token := mustToken()

	vkBot := vkAPI.New("api.vk.com", token)

	lp, _ := vkBot.DataForBotsLongPoll()

	botsLongPoll := botsLP.New(lp)

	Processor := vkBot2.NewProcessor(vkAPI.New("api.vk.com", token))

	Fetcher := vkBot2.NewFetcher(botsLongPoll)

	consumer := eventConsumer.New(Fetcher, Processor)

	if err := consumer.Start(); err != nil {
		log.Fatal("STOPPED ERROR")
	}

}

func mustToken() string {
	token := flag.String(
		"botToken",
		"",
		"write token for access vkAPI bot:")

	flag.Parse()

	if *token == "" {
		log.Fatal("bad token")
	}
	return *token
}
