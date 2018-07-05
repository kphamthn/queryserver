package main

import "regexp"

type Group struct {
	ID          string
	Name        string
	Description string
	Master      string
}

func (g Group) validate(r RequestedPlayer, old Group) (bool, string) {

	if old.Master != r.PlayerID {
		return false, "You do not have right to modify this document."
	}
	if g.Master != r.PlayerID {
		return false, "You cannot create or update this group."
	}

	if g.validateDescription() {
		return false, "Description has wrong format"
	}

	if g.validateName() {
		return false, "Name has wrong format"
	}
	return true, ""
}

func (g Group) validateDescription() bool {
	regex := "^[a-zA-Z0-9!@#$&()-`.+,/\"]{5,100}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(g.Description)
}

func (g Group) validateName() bool {
	regex := "^[a-zA-Z0-9!@#$&()-`.+,/\"]{2,50}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(g.Name)
}
