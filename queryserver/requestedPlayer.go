package main

import (
	"couchconnector"
)

type RequestedPlayer struct {
	PlayerID      string
	RequestedTime int64
}

func (r RequestedPlayer) blockedByPlayer(playerid string) bool {
	res, _ := couchconnector.GetDocumentByID(playerid)
	if !res.Get("blocked_players").Exists() {
		return false
	}
	list := res.Get("blocked_players").Array()
	for _, player := range list {
		if player.String() == r.PlayerID {
			return true
		}
	}
	return false
}

func (r RequestedPlayer) hasJoinedChallenge(challengeid string) bool {
	res, _ := couchconnector.GetViewByMultipleKeys("join", "by-challenge-player-without-reduce", []string{challengeid, r.PlayerID}, false)
	array := res.Get("rows").Array()
	if len(array) == 0 {
		return false
	}
	return true
}

func (r RequestedPlayer) bannedFromChallenge(challengeid string) bool {
	res, _ := couchconnector.GetViewByMultipleKeys("join", "by-challenge-player-without-reduce", []string{challengeid, r.PlayerID}, false)
	array := res.Get("rows").Array()
	val := array[0]
	if !val.Get("banned").Exists() {
		return false
	}
	if val.Get("banned").Bool() {
		return true
	}
	return false
}

func (r RequestedPlayer) hasJoinedGroup(groupid string) bool {
	res, _ := couchconnector.GetViewByMultipleKeys("join", "by-group-player-without-reduce", []string{groupid, r.PlayerID}, false)
	array := res.Get("rows").Array()
	if len(array) == 0 {
		return false
	}
	return true
}

func (r RequestedPlayer) bannedFromGroup(groupid string) bool {
	res, _ := couchconnector.GetViewByMultipleKeys("join", "by-group-player-without-reduce", []string{groupid, r.PlayerID}, false)
	array := res.Get("rows").Array()
	val := array[0]
	if !val.Get("banned").Exists() {
		return false
	}
	if val.Get("banned").Bool() {
		return true
	}
	return false
}
