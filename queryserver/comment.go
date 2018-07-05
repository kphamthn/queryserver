package main

import (
	"regexp"
	"time"
)

type Comment struct {
	ID          string
	Description string
	Post        string
	Player      string
}

func (c Comment) validate(r RequestedPlayer, old Comment) (bool, string) {
	b, m := c.validateUpdate(r, old)
	if !b {
		return false, m
	}
	if c.Player != r.PlayerID {
		return false, "You do  not have right to add or modify this comment."
	}
	if !c.validateContent() {
		return false, "Comment's description has wrong format"
	}
	postExists, post := getPostObject(c.Post)
	if !postExists {
		return false, "Post doesn't exist."
	}
	challengeExists, challenge := getChallengeObject(post.Challenge)
	if !challengeExists {
		return false, "Challenge doesn't exist."
	}
	if challenge.End < time.Now().UnixNano()/int64(time.Millisecond) {
		return false, "Challenge has ended."
	}
	if r.blockedByPlayer(post.Player) {
		return false, "Player is blocked by post author"
	}
	if !r.hasJoinedChallenge(challenge.ID) {
		return false, "Player hasn't joined challenge."
	}
	if r.bannedFromChallenge(challenge.ID) {
		return false, "Player is banned from challenge."
	}
	return true, ""
}

func (c Comment) validateContent() bool {
	regex := "^[a-zA-Z0-9!@#$&()-`.+,/\"]{2,500}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(c.Description)
}

func (c Comment) validateUpdate(r RequestedPlayer, old Comment) (bool, string) {
	if r.PlayerID != old.Player {
		return false, "You do not have right to modify this comment."
	}
	return true, ""
}
