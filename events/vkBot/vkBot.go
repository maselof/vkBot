package vkBot

import (
	"errors"
	"strconv"
	"vkBot/clients/botsLP"
	"vkBot/clients/vkAPI"
	"vkBot/e"
	"vkBot/events"
)

type Processor struct {
	vk *vkAPI.VkAPI
}

type Fetcher struct {
	blp *botsLP.BotsLongPollAPI
}

var ErrUnknownEventType = errors.New("unknown event type")

func NewProcessor(client *vkAPI.VkAPI) *Processor {
	return &Processor{
		vk: client,
	}
}

func NewFetcher(client *botsLP.BotsLongPollAPI) *Fetcher {
	return &Fetcher{
		blp: client,
	}
}

func (f *Fetcher) Fetch() ([]events.Event, error) {
	updates, err := f.blp.Update(f.blp.Ts)

	if err != nil {
		return nil, e.WrapIfError("can't get updates", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, upd := range updates {
		res = append(res, event(upd))
	}

	intTs, _ := strconv.Atoi(f.blp.Ts)

	intTs += 1

	f.blp.Ts = strconv.Itoa(intTs)

	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	if err := p.doCmd(event.Text, strconv.Itoa(event.UserId)); err != nil {
		return e.Wrap("can't process message", err)
	}
	return nil
}

func event(upd botsLP.Update) events.Event {
	eventType := fetchType(upd)
	eventText, eventUserId := fetchTextAndUserId(upd)

	res := events.Event{
		Type:   eventType,
		Text:   eventText,
		UserId: eventUserId,
	}

	return res
}

func fetchTextAndUserId(upd botsLP.Update) (string, int) {
	if upd.Object.Message == nil {
		return "", -1
	}

	return upd.Object.Message.Text, upd.Object.Message.FromId
}

func fetchType(upd botsLP.Update) events.Type {
	if upd.Object.Message == nil {
		return events.Unknown
	}

	return events.Message
}
