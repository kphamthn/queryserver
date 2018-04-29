package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

func validatePlayer(player gjson.Result, update bool) []byte {
	playerName := player.Get("name")
	playerMail := player.Get("email")
	playerImage := player.Get("image")
	_, err := govalidator.ValidateStruct(&Player{playerName.String(), playerMail.String(), playerImage.String()})
	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}

func validateMessage(message gjson.Result, update bool) []byte {
	messageContent := message.Get("message")
	messageType := message.Get("message_type")
	sender := message.Get("sender")
	receiver := message.Get("receiver")
	read := message.Get("read")
	_, err := govalidator.ValidateStruct(&Message{messageContent.String(), sender.String(), receiver.String(), messageType.String(), read.Bool()})
	if err != nil {
		return statusForbidden(err.Error())
	}
	return statusValid()
}

func validateChallenge(challenge gjson.Result, update bool) []byte {

	_, err := govalidator.ValidateStruct(&Challenge{challenge.Get("title").String(),
		challenge.Get("description").String(), challenge.Get("competition_mode").String(),
		challenge.Get("play_category").String(), challenge.Get("target").String(),
		challenge.Get("max_players").Int(), challenge.Get("start").Int(),
		challenge.Get("end").Int(), challenge.Get("completed").Int(),
		challenge.Get("image").String(), challenge.Get("master").String()})
	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}

func validateFriendship(friendship gjson.Result, update bool) []byte {

	_, err := govalidator.ValidateStruct(&Friendship{friendship.Get("accepted").Bool(),
		friendship.Get("player").String(), friendship.Get("friend").String()})

	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}

func validateJoin(join gjson.Result, update bool) []byte {

	_, err := govalidator.ValidateStruct(&Join{join.Get("player").String(),
		join.Get("challenge").String(), join.Get("received").Int()})

	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}

func validatePost(post gjson.Result, update bool) []byte {

	_, err := govalidator.ValidateStruct(&Post{post.Get("description").String(),
		post.Get("image").String(), post.Get("challenge").String(), post.Get("player").String()})

	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}

func validateComment(comment gjson.Result, update bool) []byte {

	_, err := govalidator.ValidateStruct(&Comment{comment.Get("description").String(),
		comment.Get("post").String(), comment.Get("challenge").String(), comment.Get("player").String()})

	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}

func validateRating(rating gjson.Result, update bool) []byte {
	_, err := govalidator.ValidateStruct(&Rating{rating.Get("player").String(),
		rating.Get("challenge").String(), rating.Get("value").Int(), rating.Get("targetID").String(),
		rating.Get("targetType").String(), rating.Get("targetPlayer").String()})

	if err != nil {
		return statusForbidden(err.Error())
	}

	return statusValid()
}
