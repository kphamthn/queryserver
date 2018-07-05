package main

import (
	"regexp"
)

type Post struct {
	ID          string
	Description string
	Image       string
	Challenge   string
	Player      string
}

func (p Post) validate(r RequestedPlayer, old Post) (bool, string) {
	if old.Player != r.PlayerID {
		return false, "You are not allowed to modify this document."
	}
	if p.Player != r.PlayerID {
		return false, "You are not allowed to post this document."
	}
	challengeExist, c := getChallengeObject(p.Challenge)
	if !challengeExist {
		return false, "Challenge does not exist."
	}
	if r.blockedByPlayer(c.Master) {
		return false, "Player blocked by challenge master."
	}
	if !r.hasJoinedChallenge(c.ID) {
		return false, "Player hasn't joined challenge."
	}
	if r.bannedFromChallenge(c.ID) {
		return false, "Player is banned from challenge."
	}
	if c.End < 0 {
		return false, "Challenge has ended."
	}
	if !p.validateContent() {
		return false, "Description has wrong format."
	}
	return true, ""
}

func (p Post) validateContent() bool {
	regex := "^[a-zA-Z0-9!@#$&()-`.+,/\"]{2,500}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(p.Description)
}
