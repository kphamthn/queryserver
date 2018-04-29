package main

import "github.com/asaskevich/govalidator"

func initCustomValidation() {
	govalidator.TagMap["messagetype"] = govalidator.Validator(func(messagetype string) bool {
		println("Jello")
		return false
	})

	govalidator.TagMap["competitionmode"] = govalidator.Validator(func(cm string) bool {
		return (cm == "pvp")
	})

	govalidator.CustomTypeTagMap.Set("numberofplayer", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case int:
			if v > 0 && v < 1000 {
				return true
			}
		}
		return false
	}))

	govalidator.CustomTypeTagMap.Set("challengedate", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := context.(type) {
		case Challenge:
			if v.End < v.Start {
				return false
			}
		}
		return false
	}))
	/*

	 */
}
