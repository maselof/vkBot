package vkBot

import (
	"log"
	"vkBot/events"
)

const (
	descriptionCmd = "Описание👾"
	filmsCmd       = "Фильмы🎬"
	serialCmd      = "Сериалы🎭"
	animeCmd       = "Аниме⛩"
	dramaCmd       = "Драма😭"
	actionCmd      = "Боевик🔫"
	comedyCmd      = "Комедия😂"
	horrorCmd      = "Хоррор😰"
	startCmd       = "Начать"
)

func (p *Processor) doCmd(text string, userId string) error {
	log.Printf("%s, %d", text, userId)

	switch text {
	case descriptionCmd, startCmd:
		return p.vk.SendMessage(msgDescription, userId, FirstKeyboard)

	case animeCmd, serialCmd, filmsCmd:
		return p.vk.SendMessage(msgWhenUserChoice, userId, SecondKeyBoard)

	case actionCmd, dramaCmd, comedyCmd, horrorCmd:
		lastMessage, _ := p.vk.GetLastMessage(userId)

		message := checkType(text, lastMessage.Response.Items[0].Text)

		return p.vk.SendMessage(message, userId, FirstKeyboard)

	default:
		return p.showKeyboard(FirstKeyboard, userId)
	}

}

func (p *Processor) showKeyboard(keyBoard events.ViewKeyBoard, userId string) error {
	err := p.vk.SendMessage(msgUnknownCmd, userId, keyBoard)

	if err != nil {
		return err
	}

	return nil
}

func checkType(t string, lastT string) string {
	switch {
	case lastT == animeCmd && t == actionCmd:
		return listAnimeAction
	case lastT == animeCmd && t == dramaCmd:
		return listAnimeDrama
	case lastT == animeCmd && t == comedyCmd:
		return listAnimeComedy
	case lastT == animeCmd && t == horrorCmd:
		return listAnimeHorror
	case lastT == filmsCmd && t == actionCmd:
		return listFilmAction
	case lastT == filmsCmd && t == dramaCmd:
		return listFilmDrama
	case lastT == filmsCmd && t == comedyCmd:
		return listFilmComedy
	case lastT == filmsCmd && t == horrorCmd:
		return listFilmHorror
	case lastT == serialCmd && t == actionCmd:
		return listSerialAction
	case lastT == serialCmd && t == dramaCmd:
		return listSerialDrama
	case lastT == serialCmd && t == comedyCmd:
		return listSerialComedy
	case lastT == serialCmd && t == horrorCmd:
		return listSerialHorror
	default:
		return msgUnknownCmd
	}
}
