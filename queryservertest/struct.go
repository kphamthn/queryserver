package main


type Player struct {
	Name  string `valid:"utfletternum, length(0|10)"`
	Email string `valid:"email"`
	Image string `valid:"url"`
}


type Message struct {
	Message     string `valid:"ascii"`
	Sender      string `valid:"playerid~Sender's ID not found \n"`
	Receiver    string `valid:"playerid~Receiver's ID not found \n"`
	MessageType string `valid:"in(image|text)~Message Type: messagetype must be text or image \n"`
	Read        bool   `valid:"-"`
}

type Challenge struct {
	Title           string `valid:"ascii"`
	Description     string `valid:"ascii"`
	CompetitionMode string `valid:"in(pvp)~Competition mode must be either pvp or \n"`
	PlayCategory    string `valid:"-"`
	Target          string `valid:"-"`
	MaxPlayer       int64  `valid:"range(1|1000)~Max Player must be a number between 1 and 1000 \n"`
	Start           int64  `valid:"challengedate~Challenge ends before it starts \n"`
	End             int64  `valid:"challengedate~Challenge ends before it starts \n"`
	Completed       int64  `valid:"-"`
	Image           string `valid:"url"`
	Master          string `valid:"playerid~Master's ID not existed or with wrong format \n"`
}

type Friendship struct {
	Accepted bool   `valid:"-"`
	Player   string `valid:"playerid~Player's ID not existed or with wrong format \n"`
	Friend   string `valid:"playerid~Friend's ID not existed or with wrong format \n"`
}

type Join struct {
	Player    string `valid:"playerid~Player'ID not existed or with wrong format \n"`
	Challenge string `valid:"challengeid~Challenge'ID not existed or with wrong format \n"`
	Received  int64  `valid:"challengedate~Challenge does not existed or is already expire \n"`
}

type Post struct {
	Description string `valid:"utfletternum"`
	Image       string `valid:"url"`
	Challenge   string `valid:"challengeid~Challenge'ID not existed or with wrong format \n"`
	Player      string `valid:"playerid~Player does not existed \n"`
}

type Comment struct {
	Description string `valid:"utfletternum"`
	Post        string `valid:"postid~Post doesn't exist"`
	Challenge   string `valid:"challengeid~Challenge'ID not existed or with wrong format \n"`
	Player      string `valid:"playerid~Player does not existed \n"`
}

type Rating struct {
	Player       string `valid:"playerid"`
	Challenge    string `valid:"challengeid~Challenge'ID not existed or with wrong format \n"`
	Value        int64  `valid:"range(-5|5)~Rating value must be a number between -5 and 5 \n"`
	TargetID     string `valid:"id~Target ID not found"`
	TargetType   string `valid:"in(comment|post|challenge)~Invalid target type"`
	TargetPlayer string `valid:"playerid~Invalid target player"`
}
