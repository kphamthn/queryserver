package main

import (
	"couchconnector"

	"github.com/tidwall/gjson"
)

func getChallengeObject(challengeID string) (bool, Challenge) {
	challenge, _ := couchconnector.GetDocumentByID(challengeID)
	if challenge.Get("error").Exists() {
		return false, Challenge{}
	}

	challengeObj := Challenge{challenge.Get("_id").String(), challenge.Get("title").String(),
		challenge.Get("description").String(), challenge.Get("competition_mode").String(),
		challenge.Get("play_category").String(), challenge.Get("target").String(),
		challenge.Get("max_players").Int(), challenge.Get("start").Int(),
		challenge.Get("end").Int(), challenge.Get("completed").Int(),
		challenge.Get("image").String(), challenge.Get("master").String()}

	return true, challengeObj
}

func getPostObject(postID string) (bool, Post) {
	post, _ := couchconnector.GetDocumentByID(postID)
	if post.Get("error").Exists() {
		return false, Post{}
	}

	postObj := Post{post.Get("_id").String(), post.Get("description").String(),
		post.Get("image").String(), post.Get("challenge").String(), post.Get("player").String()}

	return true, postObj
}

func getGroupObject(groupID string) (bool, Group) {
	group, _ := couchconnector.GetDocumentByID(groupID)
	if group.Get("error").Exists() {
		return false, Group{}
	}

	groupObj := Group{group.Get("_id").String(), group.Get("name").String(),
		group.Get("description").String(), group.Get("master").String()}

	return true, groupObj
}

/***CONVERT OBJECT***/

func getChallengeFromGjson(challenge gjson.Result) Challenge {

	challengeObj := Challenge{challenge.Get("_id").String(), challenge.Get("title").String(),
		challenge.Get("description").String(), challenge.Get("competition_mode").String(),
		challenge.Get("play_category").String(), challenge.Get("target").String(),
		challenge.Get("max_players").Int(), challenge.Get("start").Int(),
		challenge.Get("end").Int(), challenge.Get("completed").Int(),
		challenge.Get("image").String(), challenge.Get("master").String()}

	return challengeObj
}

func getPostFromGjson(post gjson.Result) Post {

	postObj := Post{post.Get("_id").String(), post.Get("description").String(),
		post.Get("image").String(), post.Get("challenge").String(), post.Get("player").String()}

	return postObj
}

func getPlayerFromGjson(player gjson.Result) Player {
	playerObj := Player{player.Get("_id").String(), player.Get("name").String(), player.Get("email").String(), player.Get("firstname").String(), player.Get("lastname").String()}
	return playerObj
}
