package main

import (
	"couchconnector"
)

type Friendship struct {
	ID       string
	Player   string
	Friend   string
	Accepted bool
}

func (f Friendship) validate(r RequestedPlayer, old Friendship) (bool, string) {
	b, m := f.validateUpdate(r, old)
	if !b {
		return false, m
	}
	if f.Player != r.PlayerID {
		return false, "Wrong playerid."
	}
	if !f.validateFriend() {
		return false, "Friend does not exist."
	}
	if old.ID == "" && f.Accepted {
		return false, "Friendship can not be accepted at the moment."
	}
	return true, ""
}

func (f Friendship) validateFriend() bool {
	res, _ := couchconnector.GetDocumentByID(f.Friend)
	if res.Get("error").Exists() {
		return false
	}
	if res.Get("type").String() != "player" {
		return false
	}
	return true
}

func (f Friendship) validateUpdate(r RequestedPlayer, old Friendship) (bool, string) {
	if old.ID != "" {
		if r.PlayerID != old.Friend && r.PlayerID != old.Player {
			return false, "You do not have the right to modify this document."
		}
		if f.Accepted != old.Accepted && f.Accepted && f.Friend != r.PlayerID {
			return false, "You are not allowed to accept this request"
		}
		if r.PlayerID != f.Friend && r.PlayerID != f.Player {
			return false, "You are not allowed to edit this document."
		}
		if f.Player != old.Player || f.Friend != old.Friend {
			return false, "Friend and player can not be changed."
		}
	}
	return true, ""
}
