package main

import "github.com/asaskevich/govalidator"

func initCustomValidation() {

	govalidator.TagMap["id"] = govalidator.Validator(func(id string) bool {
		if !regexDocumentID.MatchString(id) {
			return false
		}

		obj := getObjectByID(id)
		if !obj.Get("_id").Exists() {
			return false
		}

		return true
	})

	govalidator.TagMap["playerid"] = govalidator.Validator(func(id string) bool {
		if !regexDocumentID.MatchString(id) {
			return false
		}

		obj := getObjectByID(id)
		if obj.Get("type").String() != "player" {
			return false
		}

		return true
	})

	govalidator.TagMap["challengeid"] = govalidator.Validator(func(id string) bool {
		if !regexDocumentID.MatchString(id) {
			return false
		}

		obj := getObjectByID(id)
		if obj.Get("type").String() != "challenge" {
			return false
		}

		return true
	})

	govalidator.TagMap["postid"] = govalidator.Validator(func(id string) bool {
		if !regexDocumentID.MatchString(id) {
			return false
		}

		obj := getObjectByID(id)
		if obj.Get("type").String() != "post" {
			return false
		}

		return true
	})

	govalidator.CustomTypeTagMap.Set("challengedate", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case Challenge:
			if v.End < v.Start {
				return false
			}
		case Join:
			challenge := getObjectByID(v.Challenge)
			enddate := challenge.Get("end").Int()
			if v.Received > enddate {
				return false
			}
			return true
		}
		return true
	}))

}
