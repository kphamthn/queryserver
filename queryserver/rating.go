package main

import (
	"couchconnector"
)

type Rating struct {
	ID     string
	Player string
	Target string
	Value  int64
}

func (ra Rating) validate(r RequestedPlayer, old Rating) (bool, string) {
	if old.Player != r.PlayerID {
		return false, "You do not have right to edit this object."
	}
	if ra.Player != r.PlayerID {
		return false, "You do not have right to edit this object."
	}
	if ra.Value < -2 || ra.Value > 2 {
		return false, "Invalid Rating"
	}
	res, _ := couchconnector.GetDocumentByID(ra.Target)
	if !res.Get("type").Exists() {
		return false, "Invalid target."
	}
	if res.Get("type").String() == "challenge" {
		return ra.validateTargetChallenge(r, old, getChallengeFromGjson(res))
	}
	if res.Get("type").String() == "post" {
		return ra.validateTargetPost(r, old, getPostFromGjson(res))
	}
	if res.Get("type").String() == "player" {
		return ra.validateTargetPlayer(r, old, getPlayerFromGjson(res))
	}
	if res.Get("type").String() != "player" && res.Get("type").String() != "challenge" && res.Get("type").String() != "post" {
		return false, "Invalid Target"
	}
	return true, ""

}

func (ra Rating) validateTargetPlayer(r RequestedPlayer, old Rating, p Player) (bool, string) {
	if r.PlayerID == p.ID {
		return false, "You cannot rate yourself."
	}
	if r.blockedByPlayer(p.ID) {
		return false, "You are blocked by this player."
	}
	return true, ""
}

func (ra Rating) validateTargetChallenge(r RequestedPlayer, old Rating, c Challenge) (bool, string) {
	if r.PlayerID == c.Master {
		return false, "You cannot rate your own challenge!"
	}
	if !r.hasJoinedChallenge(c.ID) {
		return false, "You haven't join this challenge."
	}
	if r.blockedByPlayer(c.Master) {
		return false, "The challenge master has blocked you."
	}
	if r.bannedFromChallenge(c.ID) {
		return false, "You are banned from this challenge."
	}
	return true, ""
}

func (ra Rating) validateTargetPost(r RequestedPlayer, old Rating, p Post) (bool, string) {
	if r.PlayerID == p.Player {
		return false, "You cannot rate your own post."
	}
	if r.blockedByPlayer(p.Player) {
		return false, "You are blocked by the author of this post."
	}
	_, c := getChallengeObject(p.Challenge)
	return ra.validateTargetChallenge(r, old, c)
}
