package main

import "regexp"
import "couchconnector"

type Player struct {
	ID        string
	Name      string
	Email     string
	Firstname string
	Lastname  string
}

func (p Player) validation(old Player, r RequestedPlayer) (bool, string) {
	b, m := p.validateUpdate(old, r)
	if !b {
		return false, m
	}
	if !p.validateName() {
		return false, "Username doesn't have right format."
	}
	if !p.validateEmail() {
		return false, "Email doesn't have right format."
	}
	if !p.validateFirstname() {
		return false, "Firstname doesn't have right format."
	}
	if !p.checkIfEmailUnique() {
		return false, "Email must be unique."
	}
	if !p.checkIfUsernameUnique() {
		return false, "Username must be unique."
	}

	return true, ""
}

func (p Player) validateName() bool {
	regex := "^[a-zA-Z0-9._\"]{3,10}$"
	re := regexp.MustCompile(regex)
	return re.MatchString(p.Name)
}

func (p Player) validateEmail() bool {
	regex := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	re := regexp.MustCompile(regex)
	return re.MatchString(p.Email)
}

func (p Player) validateFirstname() bool {
	regex := "[a-zA-ZäöüßÄÖÜ]"
	re := regexp.MustCompile(regex)
	return re.MatchString(p.Firstname)
}

func (p Player) validateLastname() bool {
	regex := "[a-zA-ZäöüßÄÖÜ]"
	re := regexp.MustCompile(regex)
	return re.MatchString(p.Lastname)
}

func (p Player) checkIfUsernameUnique() bool {
	res, err := couchconnector.CheckIfDocumentExistWithSingleKey("player", "by-name", p.Name)
	if err != nil {
		panic(err.Error)
	}
	return res
}

func (p Player) checkIfEmailUnique() bool {
	res, err := couchconnector.CheckIfDocumentExistWithSingleKey("player", "by-email", p.Email)
	if err != nil {
		return false
	}
	return !res
}

func (p Player) validateUpdate(old Player, r RequestedPlayer) (bool, string) {
	if p.ID != r.PlayerID {
		return false, "You do not have right to change this document."
	}
	if p.Name != old.Name {
		return false, "Name cannot by changed."
	}
	return true, ""
}
