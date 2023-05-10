package eventConsumer

import (
	"log"
	"time"
	"vkBot/events"
)

type Consumer struct {
	Fetcher   events.Fetcher
	Processor events.Processor
}

func New(fetcher events.Fetcher, processor events.Processor) Consumer {
	return Consumer{
		Fetcher:   fetcher,
		Processor: processor,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.Fetcher.Fetch()
		if err != nil {
			log.Printf(err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}

	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {

		if err := c.Processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}
	return nil
}
