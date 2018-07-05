package main

import "time"

type Join struct {
	ID          string
	PlayerID    string
	ChallengeID string
	Banned      bool
	Accepted    bool
}

func (j Join) validate(r RequestedPlayer, old Join) (bool, string) {
	challengeExist, c := getChallengeObject(j.ChallengeID)
	if !challengeExist {
		return false, "Challenge does not exist."
	}
	b, m := j.validateUpdate(r, old, c)
	if !b {
		return false, m
	}
	if j.PlayerID != r.PlayerID && c.Master != r.PlayerID {
		return false, "Player does not have right to add this document."
	}
	if r.blockedByPlayer(c.Master) {
		return false, "Player is blocked by challenge master."
	}
	if old.ID != "" && old.Banned {
		return false, "Player is banned from challenge."
	}
	if old.ID == "" && r.hasJoinedChallenge(c.ID) {
		return false, "Player already has joined this challenge."
	}
	if c.End < time.Now().UnixNano()/int64(time.Millisecond) {
		return false, "Challenge has ended."
	}
	if c.getNumberOfPlayers() == c.MaxPlayer {
		return false, "Maxplayer has been reached."
	}

	return true, ""
}

/*
func (j Join) validateFields(join gjson.Result) bool {
	if !join.Get("player").Exists() ||
		!join.Get("challenge").Exists() {
		return false
	}
	return true
}
*/
func (j Join) validateUpdate(r RequestedPlayer, old Join, c Challenge) (bool, string) {
	if old.ID != "" {
		if r.PlayerID != old.PlayerID && r.PlayerID != c.Master {
			return false, "You do not have right to modify this document."
		}
		if j.PlayerID != old.PlayerID || j.ChallengeID != old.ChallengeID {
			return false, "ChallengeID and PlayerID can not be changed."
		}
		if j.Banned == true && c.Master != r.PlayerID {
			return false, "Ony challenge master can take this action."
		}
		if j.Banned != old.Banned && r.PlayerID != c.Master {
			return false, "Only challenge master can take this action."
		}
		if j.Accepted != old.Accepted && r.PlayerID != c.Master {
			return false, "Only challenge master can take this action."
		}
	}

	return true, ""
}
