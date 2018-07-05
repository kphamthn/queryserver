package main

import (
	"github.com/tidwall/gjson"
)

func validatePlayer(player gjson.Result, oplayer gjson.Result, r RequestedPlayer) []byte {
	playerObj := Player{player.Get("_id").String(), player.Get("name").String(), player.Get("email").String(), player.Get("firstname").String(), player.Get("lastname").String()}
	oldPlayerObj := Player{oplayer.Get("_id").String(), oplayer.Get("name").String(), oplayer.Get("email").String(), oplayer.Get("firstname").String(), oplayer.Get("lastname").String()}

	res, message := playerObj.validation(oldPlayerObj, r)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateChallenge(challenge gjson.Result, oldChallenge gjson.Result, r RequestedPlayer) []byte {

	challengeObj := Challenge{challenge.Get("_id").String(), challenge.Get("title").String(),
		challenge.Get("description").String(), challenge.Get("competition_mode").String(),
		challenge.Get("play_category").String(), challenge.Get("target").String(),
		challenge.Get("max_players").Int(), challenge.Get("start").Int(),
		challenge.Get("end").Int(), challenge.Get("completed").Int(),
		challenge.Get("image").String(), challenge.Get("master").String()}

	oldChallengeObj := Challenge{oldChallenge.Get("_id").String(), oldChallenge.Get("title").String(),
		oldChallenge.Get("description").String(), oldChallenge.Get("competition_mode").String(),
		oldChallenge.Get("play_category").String(), oldChallenge.Get("target").String(),
		oldChallenge.Get("max_players").Int(), oldChallenge.Get("start").Int(),
		oldChallenge.Get("end").Int(), oldChallenge.Get("completed").Int(),
		oldChallenge.Get("image").String(), oldChallenge.Get("master").String()}

	res, message := challengeObj.validation(r, oldChallengeObj)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateJoin(join gjson.Result, oldJoin gjson.Result, r RequestedPlayer) []byte {
	joinObj := Join{join.Get("_id").String(), join.Get("player").String(), join.Get("challenge").String(), join.Get("banned").Bool(), join.Get("accepted").Bool()}
	oldJoinObj := Join{oldJoin.Get("_id").String(), oldJoin.Get("player").String(), oldJoin.Get("challenge").String(), oldJoin.Get("banned").Bool(), oldJoin.Get("accepted").Bool()}
	res, message := joinObj.validate(r, oldJoinObj)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateFriendship(friendship gjson.Result, oldfriendship gjson.Result, r RequestedPlayer) []byte {
	f := Friendship{friendship.Get("_id").String(),
		friendship.Get("player").String(), friendship.Get("friend").String(), friendship.Get("accepted").Bool()}
	o := Friendship{oldfriendship.Get("_id").String(),
		oldfriendship.Get("player").String(), oldfriendship.Get("friend").String(), oldfriendship.Get("accepted").Bool()}
	res, message := f.validate(r, o)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validatePost(post gjson.Result, oldpost gjson.Result, r RequestedPlayer) []byte {
	p := Post{post.Get("_id").String(), post.Get("description").String(),
		post.Get("image").String(), post.Get("challenge").String(), post.Get("player").String()}
	o := Post{oldpost.Get("_id").String(), oldpost.Get("description").String(),
		oldpost.Get("image").String(), oldpost.Get("challenge").String(), oldpost.Get("player").String()}
	res, message := p.validate(r, o)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateComment(comment gjson.Result, ocomment gjson.Result, r RequestedPlayer) []byte {
	c := Comment{comment.Get("_id").String(), comment.Get("description").String(),
		comment.Get("post").String(), comment.Get("player").String()}
	o := Comment{ocomment.Get("_id").String(), ocomment.Get("description").String(),
		ocomment.Get("post").String(), ocomment.Get("player").String()}
	res, message := c.validate(r, o)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateRating(rating gjson.Result, orating gjson.Result, r RequestedPlayer) []byte {
	ra := Rating{rating.Get("_id").String(), rating.Get("player").String(), rating.Get("targetID").String(), rating.Get("value").Int()}
	o := Rating{orating.Get("_id").String(), orating.Get("player").String(), orating.Get("targetID").String(), orating.Get("value").Int()}

	res, message := ra.validate(r, o)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateGroup(group gjson.Result, ogroup gjson.Result, r RequestedPlayer) []byte {
	g := Group{group.Get("_id").String(), group.Get("name").String(), group.Get("description").String(), group.Get("master").String()}
	o := Group{ogroup.Get("_id").String(), ogroup.Get("name").String(), ogroup.Get("description").String(), ogroup.Get("master").String()}
	res, message := g.validate(r, o)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}

func validateJoinGroup(join gjson.Result, oldJoin gjson.Result, r RequestedPlayer) []byte {
	joinObj := JoinGroup{join.Get("_id").String(), join.Get("player").String(), join.Get("group").String(), join.Get("banned").Bool(), join.Get("accepted").Bool()}
	oldJoinObj := JoinGroup{oldJoin.Get("_id").String(), oldJoin.Get("player").String(), oldJoin.Get("group").String(), oldJoin.Get("banned").Bool(), oldJoin.Get("accepted").Bool()}
	res, message := joinObj.validate(r, oldJoinObj)
	if !res {
		return statusForbidden(message)
	}
	return statusValid()
}
