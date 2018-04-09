package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

func validatePlayer(player gjson.Result) []byte {
	playerName := player.Get("name")
	playerMail := player.Get("email")
	_, err := govalidator.ValidateStruct(&Player{playerName.String(), playerMail.String()})
	if err != nil {
		return statusForbidden(err.Error())
	}
	return statusValid()
}

func validateMessage(message gjson.Result) []byte {
	messageContent := message.Get("message")
	messageType := message.Get("messagetype")
	sender := message.Get("sender")
	receiver := message.Get("receiver")
	read := message.Get("read")
	_, err := govalidator.ValidateStruct(&Message{messageContent.String(), sender.String(), receiver.String(), messageType.String(), read.Bool()})
	if err != nil {
		return statusForbidden(err.Error())
	}
	return statusValid()
}

func validateChallenge(challenge gjson.Result) []byte {

	_, err := govalidator.ValidateStruct(&Challenge{challenge.Get("title").String(),
		challenge.Get("description").String(), challenge.Get("competitionmode").String(),
		challenge.Get("playcategory").String(), challenge.Get("target").String(),
		challenge.Get("maxplayer").Int(), challenge.Get("start").Int(),
		challenge.Get("end").Int(), challenge.Get("completed").Int(),
		challenge.Get("Image").String(), challenge.Get("Master").String()})
	if err != nil {
		return statusForbidden(err.Error())
	}
	return statusValid()
}
