package main

type JoinGroup struct {
	ID       string
	Player   string
	Group    string
	Banned   bool
	Accepted bool
}

func (j JoinGroup) validate(r RequestedPlayer, old JoinGroup) (bool, string) {
	if j.Player != r.PlayerID {
		return false, "You cannot create or update this document."
	}
	if old.ID == "" && j.Accepted {
		return false, "Request cannot accepted at the moment."
	}
	if old.ID == "" && j.Banned {
		return false, "Player cannot be banned at the moment."
	}
	groupExists, group := getGroupObject(j.Group)
	if !groupExists {
		return false, "Group doesn't exist."
	}
	b, m := j.validateUpdate(r, old, group)
	if !b {
		return false, m
	}
	if r.blockedByPlayer(group.Master) {
		return false, "Player is blocked by group master."
	}
	if r.hasJoinedGroup(group.ID) && old.ID == "" {
		return false, "You have already joined this group."
	}
	if r.bannedFromGroup(group.ID) {
		return false, "You are banned from this group."
	}
	return true, ""
}

func (j JoinGroup) validateUpdate(r RequestedPlayer, old JoinGroup, g Group) (bool, string) {
	if old.ID != "" {
		if r.PlayerID != old.Player && r.PlayerID != g.Master {
			return false, "You do not have right to modify this document."
		}
		if j.Player != old.Player || j.Group != old.Group {
			return false, "GroupID and PlayerID can not be changed."
		}
		if j.Banned == true && g.Master != r.PlayerID {
			return false, "Ony group master can take this action."
		}
		if j.Banned != old.Banned && r.PlayerID != g.Master {
			return false, "Only group master can take this action."
		}
		if j.Accepted != old.Accepted && r.PlayerID != g.Master {
			return false, "Only group master can take this action."
		}
	}

	return true, ""
}
