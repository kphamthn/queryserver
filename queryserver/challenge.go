package main

import (
	"couchconnector"
	"regexp"
	"time"
)

type Challenge struct {
	ID              string
	Title           string
	Description     string
	CompetitionMode string
	PlayCategory    string
	Target          string
	MaxPlayer       int64
	Start           int64
	End             int64
	Completed       int64
	Image           string
	Master          string
}

func (c Challenge) validation(r RequestedPlayer, old Challenge) (bool, string) {
	b, m := c.validateUpdate(r, old)
	if !b {
		return false, m
	}
	if !c.validateTitle() {
		return false, "Title has wrong format."
	}
	if !c.validateDescription() {
		return false, "Description has wrong format."
	}
	if !c.validateCompetitionMode() {
		return false, "Competition mode can be pvp or gvg."
	}
	if !c.validateTarget() {
		return false, "Target has unexpected value."
	}
	if !c.validateCategory() {
		return false, "Category has unexpected value."
	}
	if !c.validateMaxPlayer() {
		return false, "Max player must > 0 and < 10000."
	}
	if !c.validateStartEnd() {
		return false, "Challenge ends before it starts."
	}
	if !c.validateMaster(r) {
		return false, "Master is not the same as playerid."
	}

	return true, ""
}

/*
func (c Challenge) validateField(challenge gjson.Result) bool {
	checkField := false
	if !challenge.Get("title").Exists() ||
		!challenge.Get("description").Exists() ||
		!challenge.Get("competition_mode").Exists() ||
		!challenge.Get("play_category").Exists() ||
		!challenge.Get("target").Exists() ||
		!challenge.Get("max_players").Exists() ||
		!challenge.Get("start").Exists() ||
		!challenge.Get("end").Exists() ||
		!challenge.Get("completed").Exists() ||
		!challenge.Get("master").Exists() {
		return false
	}
	/*
		challenge.ForEach(func(key, value gjson.Result) bool {
			if key.String() != "title" &&
				key.String() != "description" &&
				key.String() != "competition_mode" &&
				key.String() != "play_category" &&
				key.String() != "target" &&
				key.String() != "max_players" &&
				key.String() != "start" &&
				key.String() != "end" &&
				key.String() != "completed" &&
				key.String() != "master" {
				checkField = true
				return false
			}
			return true
		})
*/
/*
		if !checkField {
			return false
		}

	return true
}
*/
func (c Challenge) validateTitle() bool {
	regex := "^[a-zA-Z0-9!@#$&()-`.+,/\"]{5,100}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(c.Title)
}

func (c Challenge) validateDescription() bool {
	regex := "^[a-zA-Z0-9!@#$&()-`.+,/\"]{2,500}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(c.Description)
}

func (c Challenge) validateCompetitionMode() bool {
	if c.CompetitionMode == "pvp" || c.CompetitionMode == "gvg" {
		return true
	}
	return false
}

func (c Challenge) validateCategory() bool {
	if c.PlayCategory == "selfChallenging" || c.PlayCategory == "openForEveryone" || c.PlayCategory == "InviteAndFriends" || c.PlayCategory == "inviteOnly" {
		return true
	}
	return false
}

func (c Challenge) validateTarget() bool {
	if c.Target == "text" || c.Target == "picture" || c.Target == "video" || c.Target == "distance" || c.Target == "weight" {
		return true
	}
	return false
}

func (c Challenge) validateMaxPlayer() bool {
	if c.MaxPlayer > 0 && c.MaxPlayer <= 10000 {
		return true
	}
	return false
}

func (c Challenge) validateStartEnd() bool {
	if c.Start < c.End {
		return true
	}
	return false
}

func (c Challenge) validateMaster(r RequestedPlayer) bool {
	if c.Master != r.PlayerID {
		return false
	}
	return true
}

func (c Challenge) validateUpdate(r RequestedPlayer, o Challenge) (bool, string) {
	if o.ID == "" {
		return true, ""
	}
	if o.Master != r.PlayerID {
		return false, "You do not have the right to modify this challenge."
	}
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if r.PlayerID != c.Master {
		return false, "Player does not have the right to change this challenge!"
	}
	if now > c.Start && c.Target != o.Target {
		return false, "Challenge's target can not be changed after it starts."
	}
	if now > c.Start && c.End != o.End {
		return false, "Challenge's deadline can not be changed after it starts."
	}
	if c.Master != o.Master {
		return false, "Challenge master can not be changed."
	}
	return true, ""
}

func (c Challenge) getNumberOfPlayers() int64 {
	result, _ := couchconnector.GetViewBySingleKey("join", "by-challenge", c.ID, true)
	array := result.Get("rows").Array()
	return array[0].Get("value").Int()
}

func (c Challenge) validateDeletion(r RequestedPlayer, old Challenge) (bool, string) {
	if r.PlayerID != old.Master {
		return false, "You do not have right to delete this."
	}
	return true, ""
}
